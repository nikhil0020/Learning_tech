## SSH

* How to connect?
    * Using password
    * using key values pair ( public and private key)
        * public key is sent to server.
        * Private key is store on client side.

* Connecting using password
    * Syntax
        * ssh UserName@SSHserver (connect to a remote host)
        * Ex - ssh root@189.110.12.10
        * This will ask for password after running this command
* known_hosts => Let the client authenticate the server to check that it isn't connecting to an impersonator.
* authorized_keys => lets the server authenticate the user.

* public key is shared with remote server.
* private key remains with client.

### Add public key to authorized_keys

* authorized_keys is the file inside .ssh directory.
* cd ~/.ssh , if it's not there , then create one 
* ssh -i .ssh/\<privateKey\> root@serverIP


### How to copy file from client to server

* scp (secure copy) => Allows you to securely copy files and directories 
* Secure meaning files and password are encrypted.
* Ex -> scp user@remotehost:/home/user/file_name
    * file_name = The name of the file that needs to be copied.
    * remoteuser =The username of the remote host.
    * remotehost = The IP address or hostname of the remote host.
    * /remote/directory = The directory where the file should be copied on the remote machine.
    * For more details - search "scp syntax" on google 
    * Ex => scp jayesh@10.143.90.2:/home/jayesh/test1.txt




