# screenshots2gif

Take screenshots and turn them into an animated GIF.

### Available flags to configure it:

| Shorthand | Flag | Type | Description | Default |
|:---------:|:-----|:----:|:------------|:-------:|
| -d | --durationSeconds | int | how many seconds long the animation should be | 3 |
| -f | --fps | int | the frame rate the animation should be made in | 10 |
| -s | --initialSleepSeconds | int | the number of seconds to wait before taking the first snapshot | 5 |
| -l | --loop | bool | if the animation should loop indefinitely | true |
| -o | --outputDir | string | the absolute path to the directory where the output gif is to be stored | project root |
| -S | --screen | int | the number identifying the screen that should be captured, default is | main screen (0) |
| -t | --timeOutMinutes | int | the number of minutes until the app is shut down as a safety measure | 5 |
| -w | --widthPixels | int | the desired width of the animation in pixels, the image will be scaled preserving aspect ratio the original value will be used if this value is set to 0 | 0 | 
  
