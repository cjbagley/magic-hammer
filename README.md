# Magic Hammer

A little helper script for my own requirements: turning images to webp, video to webm, and compressing to reduce the filesize along the way. 

Why 'Magic Hammer'? Well, it uses ImageMagick to 'hammer' stuff into shape, plus it's an item from Zelda, and I like Zelda.

## Usage:
Once compiled, add it to your $PATH and then you can run the following:

### Image Conversion
```
magic-hammer image -f [filename-to-convert]
```

| Option | Description                                                                  |
|--------|------------------------------------------------------------------------------|
| -f     | The input file to process. Default: 'input.jpg'.                             |
| -q     | The image quality value to use. Default: 82, as this gives a good balance.   |
| -tp    | The thumbnail percentage to use. If 0, it will not be resized. Default: 70%. |

### Video Conversion
```
magic-hammer video -f [filename-to-convert]
```

| Option | Description                                                                                                                      |
|--------|----------------------------------------------------------------------------------------------------------------------------------|
| -crf   | The crf value to use, from 0 to 63. The lower the number, the higher the quality (and filesize). Default: 60.                    |
| -f     | The input file to process. Default: 'input.mp4'.                                                                                 |
| -fm    | The number of minutes to start the video from. Use in conjunction with 'fs' to cut any content before the given minutes/seconds. |
| -fs    | The number of seconds to start the video from. Use in conjunction with 'fm' to cut any content before the given minutes/seconds. |
| -tm    | The number of minutes to end the video an. Use in conjunction with 'ts' to cut any content after the given minutes/seconds.      |
| -ts    | The number of seconds to end the video at. Use in conjunction with 'tm' to cut any content after the given minutes/seconds.      |

Note regarding the times: if you set both a start and an end time, the end time is from the new start time.
For example, if you have a 20 second clip, you want to start 5 seconds in and cut the last 5 seconds, you will need:
-fs 5 -ts 10
Basically, you are saying: start 5 seconds in, and stop after 10 seconds (which would leave you 5 seconds from the end).
