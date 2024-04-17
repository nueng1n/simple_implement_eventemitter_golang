Simple Implementation of Create EventEmitter on golang like nodejs.

Nodejs
```
   const eventEmitter = new EventEmitter();

    eventEmitter.on('example1', (msg) => {
        console.log("Received message  on `example1` event", msg);
    });

    eventEmitter.on('example2', (msg) => {
        console.log("Received message  on `example2` event", msg);
    });

    eventEmitter.emit('example1', 'send a');
   
    eventEmitter.emit('example2', 'send b');
    eventEmitter.emit('example1', 'send c');
```

Golang

```
  ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	emitter := NewEventEmitter(ctx, &wg)

	emitter.On("example1", func(str string) {
		log.Println("Received message  on `example1` event", str)
	})

	emitter.On("example2", func(str string) {
		log.Println("Received message  on `example2` event", str)
	})

	wg.Add(1)
	emitter.Emit("example1", "send a")

	wg.Add(1)
	emitter.Emit("example2", "send b")
	wg.Add(1)
	emitter.Emit("example1", "send c")

	wg.Wait()
```
