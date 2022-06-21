# https://github.com/ArduCAM/mjpg-streamer/blob/master/mjpg-streamer-experimental/plugins/output_http/README.md
MJPG_DIR=/home/pi/Projects/mjpg-streamer/mjpg-streamer-experimental
$MJPG_DIR/mjpg_streamer -o "output_http.so -l 0.0.0.0 -p 8081" -i "input_libcamera.so"

echo "http://127.0.0.1:8080/?action=stream"
# ./mjpg_streamer -o "output_http.so -w ./www" -i "input_libcamera.so"

#  Help for input plugin..: Libcamera Input plugin
#  ---------------------------------------------------------------
#  The following parameters can be passed to this plugin:

#  [-r | --resolution ]...: the resolution of the video device,
#                           can be one of the following strings:
#                           QQVGA QCIF CGA QVGA CIF PAL 
#                           VGA SVGA XGA HD SXGA UXGA 
#                           FHD 
#                           or a custom value like the following
#                           example: 640x480
#  [-f | --fps ]..........: frames per second
#  [-b | --buffercount ]...: Set the number of request buffers.
#  [-q | --quality ] .....: set quality of JPEG encoding
#  ---------------------------------------------------------------
#  Optional parameters (may not be supported by all cameras):

#  [-br ].................: Set image brightness (integer)
#  [-co ].................: Set image contrast (integer)
#  [-sa ].................: Set image saturation (integer)
#  [-ex ].................: Set exposure (integer)
#  [-gain ]...............: Set gain (integer)
#  ---------------------------------------------------------------