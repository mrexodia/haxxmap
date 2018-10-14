# haxxmap

Some simple go tools to perform a Man-in-the-middle (MITM) attack on your IMAP server in case you forgot your password.

## Use case

I forgot the password to my email account, but on my iPhone Mail was still working fine. The idea is to proxy the IMAP server and retrieve the password from there.

### Project structure

`client` is an example client for local testing (just displays the subjects of the 4 latest messages).

`dns` contains the [CoreDNS](https://github.com/coredns/coredns) configuration file to redirect `imap.example.com` to an IP in the local network.

`proxy` contains the IMAP proxy used to retreive the password.

`server` contains a test server.

### Setting up

The first step is to create a self-signed certificate for `imap.example.com` and add it as a trusted root on the iPhone. For this purpose a website like http://www.selfsignedcertificate.com will do! To add a certificate you have to send it to your phone (with AirDrop for example), then click `Install`. After that go to Settings -> General -> About -> Certificate Trust Settings and flick the checkbox next to `imap.example.com`.

Next up is configuring the DNS. You can use `sudo coredns` from the `dns` folder and everything should work (you might have to change the local IP and domain to match your details). Then manually configure DNS on the phone to point to your server (DNS requests to other domains will be resolved/cached normally).

Now just do `sudo go run proxy.go imap.example.com:993 cert.pem cert.key` from the `proxy` folder to start the IMAP server.

### Getting the password

Simply go to the Mail app and refresh. You should see the password in the IMAP proxy console!