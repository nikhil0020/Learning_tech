## FFMPEG

#### Converting file format using ffmpeg

- **ffmpeg -i input.avi output.mp4**
- Here we are converting .avi format to .mp4
- we can also convert .mp4 to .mp4 -> video to audio
- -i here means input

### File quality

- #### For .avi format

  - We have deal will quantizer , lower the quantizer higher the quality and vice versa. higher the quality, higher the size.
  - **ffmpeg -i input.avi -q 23 output.avi**
  - number between [20,30] is common but if we want higher quality reduce the number and for lower quality increase the number.

- #### For .mp4 format
  - We have to use **-crf** value
  - **ffmpeg -i input.mp4 -crf 18 output.mp4**

### Bit rate

- #### for audio channel
  - **ffmpeg -i input.mp3 -b:a 320k output.mp3**
  - -b indicates bitrate and -b:a indicates bitrate for audio
  - <a href="https://restream.io/blog/what-is-video-bitrate"/>How bitrate affect video quality, and other factor that affect video quality</a>
