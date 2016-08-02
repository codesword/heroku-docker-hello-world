FROM alpine:latest

# Install bash, It's needed by heroku
RUN apk add --update bash

# Copy the generated static binaries
COPY artifact /

ENV PORT 8000
# Expose is NOT supported by Heroku
# EXPOSE 50050
EXPOSE 8000
# Run the app.  CMD is required to run on Heroku
# $PORT is set by Heroku
CMD ["./artifact"]
