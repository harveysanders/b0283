libcamera-vid \
--verbose
--timeout 0 \
--inline \
--output - | cvlc -vvv stream:///dev/stdin \
--sout '#rtp{sdp=rtsp://:8080/stream1}' \
:demux=h264

# echo "Open URL in VLC: rtsp://derbypi.local:8080/stream1"