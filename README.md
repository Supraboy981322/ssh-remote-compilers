# Remote compilation over SSH

## Why?

For some reason, binaries that I compile on my PC (x86 NixOS) don't run on any other x86 Linux machine, but binaries compiled on other machines run fine on both it and any other x86 Linux machine. So in order to have portable binaries, I must compile on a separate machine. I was doing it semi-manually over SSH, but I got tired of writing the following command (for example)
```sh
cd ..; tar -cf - src/* | ssh super@dev 'set -e; mkdir -p /tmp/foo; cd /tmp/foo; tar -xf -; cd src; go build -o out.bin .;cat out.bin; cd /tmp; rm -r foo' > foo; cd src
```
This project is a server and client that does this automatically

## Usage

The client isn't written yet, but you can still use it over SSH with a command like this:

(replace `your-server` with the address for your SSH server and `foo` with your binary name)
```sh
tar -cf - . | ssh go@your-server -p 7845 . > foo
```
  You still have to make it executable (you can add `sudo chmod a+x foo` to the end of the chain for this).

Once the client is written, usage will be more like this:
```sh
rC go -o foo .
```
