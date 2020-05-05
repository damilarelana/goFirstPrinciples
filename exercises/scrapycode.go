
/* parsePlotData()
 - takes in the returned plotDataMap
 - parses it through a for loop so as to generate a `array of array`
   + to extract the [index, stateData] at each dictionary key
*/
// func parsePlotData(plotDataDict map[int][]int) {
// 	mapLength := len(plotDataDict)  // get number of items in the map i.e. how many stateData do we have
// 	stateDataArrays := make([]int, mapLength)
// 	for
// 		stateDataArrays := [values for values in plotDataDict.values()]
// 	return stateDataArrays  // returns the list of lists required by animation
// }


/*
 createAnimation()
   - takes in the stateDataLists
   - takes in listMinValue
   - takes in listMaxValue
   - takes in the algorithm's name e.g. "insertion Sort" etc.
   - takes in the animation file format e.g. "mp4" or "gif"
   - uses the parameters to create an animation of the sorting
*/
func createAnimation(stateDataLists []int, listMinValue int[], listMaxValue []int, algorithmName string, animationFormat string, speedFlag bool){

  // determine length of each stateData list (in the list of list) meant to be animated
  inputListLength, ok = len(stateDataLists[0])
  if !ok {
    ErrMsgHandler("Unable to determine length of stateData list (in the `array of arrays`)", nil)
  }

  /* setup the mesh grid
     - xLinspace and yLinspace input to np.meshgrid()
     - xx and yy output indicating the width and height of the mesh grid in terms of array
     - initializing the meshGrid at the beginning would not be a problem even if stateData length changes during sorting
         + this is because the stateData length can only reduce and not increase
         + thus the meshGrid would still be big enough to handle rendering on it by smaller data set
  */
  xLinspace = np.linspace(0, inputListLength-1, inputListLength)  // creates i.e. evenly spaced stuff in x-axis that matches the array index spacing
  yLinspace = np.linspace(listMinValue, listMaxValue, inputListLength) // creates even space in y-axis for array element values
  xx, yy = np.meshgrid(xLinspace, yLinspace)  // create the mesh grid 

  // setup the matplotlib's plot parameters
  fig, ax = setupPlotParams(listMinValue, listMaxValue, inputListLength)
  fig.show()
  // fig.canvas.draw()

  /* create animation
     - 'interval' talks about ms interval between frames
     - 'frames' required to help save the animation later
     - 'blit' ensures that only areas of the plot which have changed are re-drawn i.e. improves performance and smoothness
     - uses `fig` after setupPlotParams() is called
     - calls `animate()` [while using the `stateDataList` and `arrayPlot`] defined within the scope of `createAnimation`
  */

  /* numEvents
     - refers to the number of stateData that exists in the stateDataLists
     - i.e. how many lists do you have within the `list of lists`
     - where each `list` represents an event
  */
  numEvents = len(stateDataLists)

  /* plot the first array data
     - note that `stateDataLists[0, :]` is acting like `yy` i.e. the height data to the barplot handler
     - hence why we later `animate` i.e. iterate of `yy` (i.e. stateDataLists[i, :] different indices) in the `animate()` function
     - assumption is that:
         + length of xx[i] does not change AND length of stateDataList[i] also does not change
         + hence this breaks with mergeSort() since stateDataList[i] changes during mergeSort()
         + two solutions:
             - Option I: dynamically alter length of `xx[0]` i.e. recomputing np.meshgrid everytime [with new Linspace], while stateData[i] changes
             - Option II: `xx[0]` remains the same but stateData[i] is dynamically padded (at the) with `zeros` to simulate absence of plot data [as mergeSort splits the data down to only one index]
         + Option is more DRY and SOLID i.e. it avoids to many changes to the animation convas
             - only the plotting data is dynamically changing
  */
  arrayPlot = plt.bar(xx[0], stateDataLists[0], 0.8, None, color='green', edgecolor='snow', alpha=0.7)  // helps to ensure that we plot only the first array due to how ax.bar handles 
  
  backgroundRender = fig.canvas.copy_from_bbox(ax.bbox)  // save background be fore animating ... to improve performance
  animStartTime = time.time()
  animate(xx[0], ax, arrayPlot, stateDataLists, fig, numEvents, backgroundRender, speedFlag)
  print('animation rendered in {:.2f}s'.format(time.time()-animStartTime))  // print the fps to have 2 decimal places

  /* save animation
     - checks the animation format
     - uses anonymous function to simulate a switch statement
     - to save as `mp4` change `filename` assignment i.e. "filename = algorithm + '.mp4' "
  */
  filename = algorithmName + animationFormat
  // if animationFormat == "gif":
  //     animation.save(filename, writer='imagemagick')
  // else:
  //     animation.save(filename)
}


// animate()
func animate(x, ax, arrayPlot, stateDataLists, fig, numEvents, backgroundRender, speedFlag){
  for i,_ := range(numEvents){
      fig.canvas.restore_region(backgroundRender) // restore the cached/rendered background (i.e. apart from the bars themselves)
      arrayPlot.remove() // remove the previous bar chart rendering, so as to avoid ghosting

      if hasattr(arrayPlot, 'set_height'){
          arrayPlot.set_height(stateDataLists[i])  // faster option but does not always work due to matplotlib issues
      } else {
        arrayPlot = plt.bar(x, stateDataLists[i], 0.8, None, color='green', edgecolor='snow', alpha=0.7)  // brute force creation of new bar container
        if speedFlag is true { 
          time.sleep(0.5)  // meant to slow down the animation plot of mergesort() data
          fig.canvas.draw()
          fig.canvas.blit(ax.bbox)
          fig.canvas.flush_events()
        } else{
          fig.canvas.draw()
          fig.canvas.blit(ax.bbox)
          fig.canvas.flush_events()
        }
      }
  }
}

// setupPlotParameters()
// - sets up the required initial plot parameters
func setupPlotParams(arrayMinValue int, arrayMaxValue int, inputArrayLength int){

  /* plt.ion()
   setup plotting area attributes
     - set_ylim refers to the the actual values in the array hence why it should be on the y-axis - in relation to a bar chart
         + `5` is used to pad the listMinValue and listMaxValue to animating those values easier to see
     - set_xlim refers to the index numbers of the array, which can be computed simply from the array length 
     - xLinspace is obtained from the global variable scope
     - yLinspace is calculated here using attributes of the already sorted list e.g. listMaxValue
   figsize = (6, 3)  # set figure size tuple i.e. canvas size (i.e. paper size: A4, Letter, wide-screen aspect ratio etc.)
   ax.set(xlim=(0, inputListLength), ylim=(listMinValue-5, listMaxValue+5))
     - note that xlim(0, inputListLength) was choosen as such to ensure that last element is visible during the plot
     - it would have been invisible, if xlim(0, inputListLength) was used instead
  */

  fig = plt.figure(figsize=(10, 6))
  ax = fig.add_subplot(1,1,1, xlim = (-1, inputListLength), ylim = (listMinValue-20, listMaxValue+20), xlabel = "Index", ylabel = "Value", title = algorithmName, alpha=0.6)
  
  return fig, ax
}



// createAnimation()
func createAnimation(plotDataMap map[int][]int, arrayLength int, algorithmName string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			lastIndex := arrayLength - 1
			xAxisItems := createRandomArray(0, lastIndex, 1)
			bar := charts.NewBar()
			bar.SetGlobalOptions(charts.TitleOpts{Title: algorithmName})
			bar.AddXAxis(xAxisItems)
			f, err := os.Create("bar.html")
			for index, stateData := range plotDataMap {
				bar.AddYAxis("Array State", stateData)
				if err != nil {
					errMsg := fmt.Sprintf("Unable to create bar.html for stateData at plotDataMap index %d", index)
					ErrMsgHandler(errMsg, err)
				}
				bar.Render(w, f)
			}
		})
}


func baseBar(xAxisItems []int, algorithmName string, plotDataArrays [][]int) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.TitleOpts{Title: algorithmName},
		charts.ToolboxOpts{Show: true},
	)
	bar.AddXAxis(xAxisItems).AddYAxis("Array State", plotDataArrays)
	return bar
}

// createAnimation()
func createAnimation(plotDataArrays [][]int, arrayLength int, algorithmName string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			lastIndex := arrayLength - 1
			xAxisItems := createRandomArray(0, lastIndex, 1)
			page := charts.NewPage()
			page.Add(
				baseBar(xAxisItems, algorithmName, plotDataArrays),
			)
			f, err := os.Create("bar.html")
			if err != nil {
				errMsg := fmt.Sprintf("Unable to create bar.html for plotDataArray")
				ErrMsgHandler(errMsg, err)
			}
			page.Render(w, f)
		})
}

// baseBar
func baseBar(xAxisItems []int, algorithmName string, plotDataArrays [][]int) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.TitleOpts{Title: algorithmName},
		charts.ToolboxOpts{Show: true},
	)
	for _, stateData := range plotDataArrays {
		bar.AddXAxis(xAxisItems).AddYAxis("Array State", stateData)
	}
	return bar
}


// createAnimation()
func createAnimation(plotDataArrays [][]int, arrayLength int, algorithmName string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			lastIndex := arrayLength - 1
			xAxisItems := createRandomArray(0, lastIndex, 1)
			page := charts.NewPage()
			for _, stateData := range plotDataArrays {
				page.Add(
					baseBar(xAxisItems, algorithmName, stateData),
				)
				f, err := os.Create("bar.html")
				if err != nil {
					errMsg := fmt.Sprintf("Unable to create bar.html for plotDataArray")
					ErrMsgHandler(errMsg, err)
				}
				page.Render(w, f)
			}
    })
    


// createAnimation()
func createAnimation(plotDataArrays [][]int, arrayLength int, algorithmName string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			lastIndex := arrayLength - 1
			xAxisItems := createRandomArray(0, lastIndex, 1)
			page := charts.NewPage()
			for index, stateData := range plotDataArrays {
				if index != 0 {
					err := os.Remove("bar.html")
					if err != nil {
						errMsg := fmt.Sprintf("Unable to remove the 'bar.html' for stateData %d, before new render for stateData %d", index-1, index)
						ErrMsgHandler(errMsg, err)
					}
				}
				f, err := os.Create("bar.html")
				if err != nil {
					errMsg := fmt.Sprintf("Unable to create bar.html for plotDataArray")
					ErrMsgHandler(errMsg, err)
				}
				page.Add(
					baseBar(xAxisItems, algorithmName, stateData),
				)
				page.Render(w, f)
			}
		})
}



// baseBar
func baseBar(xAxisItems []int, algorithmName string, stateData []int) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.TitleOpts{Title: algorithmName},
		charts.ToolboxOpts{Show: true},
	)
	bar.AddXAxis(xAxisItems).AddYAxis("Array State", stateData)
	return bar
}

// baseBar
func baseBar(xAxisItems []int, algorithmName string, plotDataArrays [][]int) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.TitleOpts{Title: algorithmName},
		charts.ToolboxOpts{Show: true},
	)
	for index, stateData := range plotDataArrays {
		if index == 0 {
			bar.AddXAxis(xAxisItems).AddYAxis("Array State", stateData)
		} else {
			bar.Overlap(bar.AddXAxis(xAxisItems).AddYAxis("Array State", stateData))
		}
	}
	return bar
}