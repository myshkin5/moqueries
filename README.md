# Moqueries - Lock-free interface and function mocks for Go
Moqueries makes mocks, but not just interface mocks &mdash; `moqueries` builds mocks for functions too. But these aren't your typical mocks!

Moqueries mocks are, as mentioned above, lock free. Why would that matter? A single lock per mock shouldn't slow down your unit tests that much, right? The problem isn't speed but semantics. Having all interactions in your tests synchronized by locking just to protect mock state changes the nature of the test. That lock in your old mock could be hiding subtle bugs from Go's race detector!

These mocks are also true to the interface and function types they mock &mdash; several mock generators record your intentions with method signatures like `DoIt(arg0, arg1, arg2 interface{})` when your interface is something like `DoIt(lFac, rFac *xyz.Factor, msg string)`. This applies to both parameters passed into the recorder and result values. Having a true implementation means that your IDE and the compiler both know what the types should be which improves your cycle time.

## Generating mocks
Mocks are generated by directly invoking `moqueries` from your terminal but typically you want to capture the command in a `//go:generate` directive. The following directive (in a working example [here](https://github.com/myshkin5/moqueries/blob/master/demo/demo.go#L10)) generates a mock for Go's ubiquitous [`io.Writer`](https://golang.org/pkg/io/#Writer) interface:
```go
//go:generate moqueries --destination moq_writer_test.go --import io Writer
```

Note that because of the `--destination` option, the mock is written to a file called [`moq_writer_test.go`](https://github.com/myshkin5/moqueries/blob/master/demo/moq_writer_test.go) in the same directory containing this directive. Also note that since we presumably didn't put this directive in Go's standard library `io` package, we have to include a `--import io` option so that Moqueries can find the type.

Generating mocks for function types is just as easy. In the same example ([a few lines further down](https://github.com/myshkin5/moqueries/blob/master/demo/demo.go#L12-L14)), we put a Go generate directive directly above the type definition (the best place for the directive unless you are mocking third-party types):
```go
//go:generate moqueries --destination moq_isfavorite_test.go IsFavorite

type IsFavorite func(n int) bool
```

## Using mocks

### Creating a mock instance
Code generation creates a `newMockXXX` factory function for each mock you generate. Simply [invoke the function and hold on to the new mock](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L17-L18) for further testing:
```go
isFavMock := newMockIsFavorite(scene, nil)
writerMock := newMockWriter(scene, nil)
```

You might be curious what that `scene` parameter is. A scene provides an abstraction on a collection of mocks. It allows your tests to control all of its mocks at once. There are more details on the use of scenes [below](#working-with-multiple-mocks) but for now, you can create a scene like [this](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L16):
```go
scene := moq.NewScene(t)
```

### Expectations
To get a mock to perform specific behaviors, you have to tell it what to expect and how to behave. For function mocks, the `onCall` function (generated for you) has the same parameter signature as the function itself. The return value of the `onCall` function is a type that (via its `returnResults` method) informs the mock what to return when invoked with the given set of parameters. For our `IsFavorite` function mock, we tell it to expect to be called with parameters `1`, `2` and then `3` but only `3` is our favorite number [like so](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L20-L22):
```go
isFavMock.onCall(1).returnResults(false)
isFavMock.onCall(2).returnResults(false)
isFavMock.onCall(3).returnResults(true)
```

Working with interface mocks is very similar to working with function mocks. For interface mocks, the generated `onCall` method returns the expectation recorder of the mocked interface (a full implementation of the interface for recording expectations). For our `Writer` mock example, we tell it to expect a call to `Write` with the [following code](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L24-L25):
```go
writerMock.onCall().Write([]byte("3")).
    returnResults(1, nil)
```

Note in the above code, we told the mock to return `1` and `nil` with a call to the generated `returnResults` method. Per the interface definition of a writer, we wrote one byte successfully with no errors. Alternatively, we could indicate an error with [a small change](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L48-L49):
```go
writerMock.onCall().Write([]byte("3")).
    returnResults(0, errors.New("couldn't write"))
```

### Arbitrary (any) parameters
Sometimes it's hard to know what exactly the parameter values will be when setting expectations. You want to say "ignore this parameter" when setting some expectations. The generated `anyXXX` functions do exactly that &mdash; the specified parameter will be ignored (in the recorded expectation and again later when the mock is invoked). The [following code](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L162-L163) sets the expectation for a function called `GadgetsByWidgetId` that takes a single `int` parameter called `widgetId`. With this expectation, the mock will return the same result regardless of the value of `widgetId`:
```go
storeMock.onCall().GadgetsByWidgetId(0).anyWidgetId().
    returnResults(nil, nil).times(2)
```

Expectations with more match parameters are given precedence over expectations with fewer matching parameters. In another test, we work with another mocked method called `LightGadgetsByWidgetId` that presumably returns gadgets associated with a specified widget that are less than a specified weight. The [following snippet](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L207-L208) returns the `g1` and `g2` gadgets when `LightGadgetsByWidgetId` is called with a `widgetId` of `42` regardless of the value specified for `maxWeight`:
```go
storeMock.onCall().LightGadgetsByWidgetId(42, 0).anyMaxWeight().
    returnResults([]demo.Gadget{g1, g2}, nil)
```

In the same test, [these lines](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L221-L222) return `g3` and `g4` regardless of either parameter specified to `LightGadgetsByWidgetId`:
```go
storeMock.onCall().LightGadgetsByWidgetId(0, 0).anyWidgetId().anyMaxWeight().
    returnResults([]demo.Gadget{g3, g4}, nil)
```

Callers will be returned `g3` and `g4` unless the `widgetId` is `42`, in which case they will be returned `g1` and `g2`.

### N-times and Any times
When expectations need to be returned repeatedly, `times` and `anyTimes` can be used to control how often a particular result is returned. For instance, [the following code](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L74-L76) instructs the mock function to return `false` five times and then `true` one time (one time is the default):
```go
isFavMock.onCall(7).
    returnResults(false).times(5).
    returnResults(true)
```

`anyTimes` instructs the mock to repeatedly return the same values regardless of how many times the function is called with the given parameters. Note that `anyTimes` can only be used once for a given set of parameters. In fact, `anyTimes` has no return value so no other expectations can be set after it.

`times` and `anyTimes` can be used together as well. [This code](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L133-L135) returns `true` twice and then always returns `false` regardless of how many times the function is called with the parameter `7`:
```go
isFavMock.onCall(7).
    returnResults(true).times(2).
    returnResults(false).anyTimes()
```

### Asserting call sequences
Some test writers want to make sure not only were certain calls made but also that the calls were made in an exact order. If you want to assert that all calls for a test are to be in order, just set the mock's [default to use sequences](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L98) on all calls via the `MockConfig` value:
```go
config := moq.MockConfig{Sequence: moq.SeqDefaultOn}
```

Now the calls to all mocks using the above config must be in the exact sequence. The sequence of expectations must match the sequence of calls to the mock.

If there are just a few calls that must be in a specific order relative to each other, [call the `seq` method](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L263-L265) when recording expectations:
```go
isFavMock.onCall(1).seq().returnResults(false)
isFavMock.onCall(2).seq().returnResults(false)
isFavMock.onCall(3).seq().returnResults(true)
```

This is basically overriding the default so that just the calls specified use a sequence. You can also turn off sequences on a per-call basis when the default is to use sequences on all calls [using the `noSeq` method](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L293-L294):
```go
writerMock.onCall().Write([]byte("3")).noSeq().
    returnResults(1, nil)
```

### Do functions
Sometimes you need to tap into what your mock is doing. You may need to capture a value that was passed to a mock or you may need to have some logic calculate what a mock's response should be. Do functions do just that. If you just need to listen in to a `returnResults` expectation, you provide a [function that matches the mocked functions parameters](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L314-L317) (in this case the mocked function takes a single `int` parameter):
```go
sum := 0
sumFn := func(n int) {
    sum += n
}
```

Then [chain](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L319-L321) an `andDo` call after the `returnResults` call:
```go
isFavMock.onCall(1).returnResults(false).andDo(sumFn)
isFavMock.onCall(2).returnResults(false).andDo(sumFn)
isFavMock.onCall(3).returnResults(true).andDo(sumFn)
```

If on the other hand you need to calculate a mock's return values, start with [a function that has the same signature as the mocked function](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L348-L350) (both parameters and result values):
```go
isFavFn := func(n int) bool {
    return n%2 == 0
}
```

Now you can replace both the `returnResults` and `andDo` calls with [a single call](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L352) to `doReturnResults`:
```go
isFavMock.onCall(0).anyN().doReturnResults(isFavFn).anyTimes()
```

Note this expectation will return the calculated value (`n%2 == 0`) regardless of the input parameters and regardless of how may times it is invoked.

### Passing the mock to production code
Each mock gets a generated `mock` method. This function accesses the implementation of the interface or function invoked by production code. In [our example](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L27-L30), we have a type called `FavWriter` that needs an `IsFavorite` function and a `Writer`:
```go
d := demo.FavWriter{
    IsFav: isFavMock.mock(),
    W:     writerMock.mock(),
}
```

### Nice vs. Strict
Sometimes your mocks will get lots of function calls with lots of different parameters &mdash; maybe more calls than you can (or want) to configure. Nice mocks trigger special logic that allow them to return zero values for any unexpected calls. [Creating a nice mock](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L42-L43) is as simple as supplying a little configuration to the `new` method (the value was `nil` above which defaults to creating strict mocks):
```go
isFavMock := newMockIsFavorite(
    scene, &moq.MockConfig{Expectation: moq.Nice})
```

Now we only need to [set expectations](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L46) when returning non-zero values (`returnResults(false)` is now the default):
```go
isFavMock.onCall(3).returnResults(true)
```

Calling this mock with any value besides `3` will now return `false` (without having to register any other expectations).

### Asserting all expectations are met
After your test runs, you may want to verify that all of your expectations were actually met. Each mock implements `AssertExpectationsMet` to [do just that](https://github.com/myshkin5/moqueries/blob/master/demo/demo_test.go#L37):
```go
writerMock.AssertExpectationsMet()
```

### Resetting the state
Occasionally you need reset the state of a mock. Re-creating the mock is preferred but there are situations where that isn't possible (maybe a long-running test, or the mock has already been handed off to other code). In any case, calling `Reset` does just that &mdash; it resets the mock:
```go
writerMock.Reset()
```

### Working with multiple mocks
Quite often tests will require several mocks. A `Scene` is a collection of mocks, and it allows you to perform actions on all the mocks with a single call. Both `AssertExpectationsMet` and `Reset` are supported:
```go
scene.AssertExpectationsMet()
```

## More command line options
Below is a loose collection of out-of-the-ordinary command line options for use in out-of-the-ordinary situations.

### Mocking interfaces and function types in test packages
When the type you want to mock is defined in a test package, use one of the following two solutions:

1. Specify the test package (with its `_test` suffix) in the `--import` option:
    ```go
    //go:generate moqueries --destination moq_isfavorite_test.go --import github.com/myshkin5/moqueries/demo_test IsFavorite
    ```
   Note: This solution requires the `--import` option even if your Go generate directive is in the same package being imported.

   *_&mdash; OR &mdash;_*

1. Use the `--test-import` option:
    ```go
    //go:generate moqueries --destination moq_isfavorite_test.go --test-import IsFavorite
    ```

### Exported (public) mocks
Mocks are typically generated in the test package of the destination directory. This works well in most cases including when the code you want to test lives in the same package as the code you wan to mock out. When you have lots of different packages all using the same mocks, it's better to generate the mocks once and import the mocks where needed. This is where the `--export` command line option comes into play:
```go
//go:generate moqueries --destination moq_writer.go --export --import io Writer
```

Now all of the mock's structs and methods are exported, so they can be used from any package:

```go
writerMock := mockpkg.NewMockWriter()

writerMock.OnCall().Write([]byte("3")).ReturnResults(0, errors.New("couldn't write"))
```

## Command line reference
The Moqueries command line has the following form:

```bash
$ moqueries [options] [interfaces and/or function types to mock] [options]
```

Interfaces and function types are separated by whitespace. Multiple types may be specified.

| Option | Type | Default value | Usage |
|---|---|---|---|
| `--debug` | `bool` | `false` | If true, debugging output will be logged |
| `--destination <file>` | `string` | None, required | The file path where mocks are generated relative to directory containing generate directive (or relative to the current directory) |
| `--export` | `bool` | `false` | If true, generated mocks will be exported and accessible from other packages |
| `--import <name>` | `string` | `.` (the directory containing generate directive) | The package containing the type (interface or function type) to be mocked |
| `--package <name>` | `string` | The test package of the destination directory when `--export=false` or the package of the destination directory when `--export=true` | The package to generate code into |
| `--test-import` | `bool` | `false` | Directs Moqueries to look for the types to be mocked in the test package |

Values specified after an option are separated from the option by whitespace (`--destination moq_isfavorite_test.go`) or by an equal sign (`--destination=moq_isfavorite_test.go`).

If a non-repeating option is specified more than one time, the last value is used.

Options with a value type of `bool` are set (turned on) by specifying the option directly (`--debug`) or by specifying a boolean value after the option (`--debug=true` or `--debug true`). To turn a boolean option off, a value must be specified (`--debug=false`).
