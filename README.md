# Remote compilation over SSH

## Why?

For some reason, binaries that I compile on my PC (x86) don't run on any other x86 machine, but binaries compiled on other machines run fine on both it and any other x86 machine. So in order to have portable binaries, I must compile on a separate machine. I was doing it semi-manually over SSH, but I got tired of writing the following command (for example)
```sh
cd ..; tar -cf - src/* | ssh super@dev 'set -e; mkdir -p /tmp/foo; cd /tmp/foo; tar -xf -; cd src; go build -o out.bin .;cat out.bin; cd /tmp; rm -r foo' > foo; cd src
```
This project is a server and client that does this automatically
