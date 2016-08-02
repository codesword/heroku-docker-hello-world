FROM alpine:latest

# Install bash, It's needed by heroku
RUN apk add --update bash

# Copy the generated static binaries
COPY artifact /

# Expose is NOT supported by Heroku
# EXPOSE 80

# Run the app.  CMD is required to run on Heroku
# $PORT is set by Heroku
CMD ["./artifact"]
