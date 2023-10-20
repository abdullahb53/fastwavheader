# FastWavHeader
 __Optimized way to get wav header.__<br/>

![fastwavheaderAlloc](https://github.com/abdullahb53/fastwavheader/assets/29378922/8e63ad06-eea0-4daf-a9ca-02bd4254670f)


**`Slice/Unsafe is selected`**

 &nbsp;<sub>Show becnhmark results:<sub>
```
make test
```

```golang
	// If u want to stream data through channels.
	fwh.StartStreamEvent()
	filePaths := []string{
         ...
	}

	filePaths2 := []string{
         ...
	}

	// Send your filePaths to channel.
	for _, val := range filePaths {
		fwh.FilePathCh <- val
	}

	// Consume WavHeaderInfos from channel.
	go func() {
		consume := fwh.HeaderCh
		for {
			select {
			case wavinfo, ok := <-consume:
				if !ok {
					consume = fwh.HeaderCh
				} else {
                                        // Do something..
					log.Println("@@ WavInfo:", wavinfo)
				}
			default:
				continue
			}
		}
	}()

	// Change channel capacity. Async or Sync.
	go fwh.ChangeQueueSize(30, 40)

	// Again, send new filePaths to adjusted channel.
	for _, val := range filePaths2 {
		fwh.FilePathCh <- val
	}
```
