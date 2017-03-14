# Sudo Shell

Sudo Shell is a wrapper to run a login shell with `sudo` for the purpose of session audit logging. 

[![Build Status](https://travis-ci.org/cloudposse/sudosh.svg?branch=master)](https://travis-ci.org/cloudposse/sudosh)
[![GitHub Stars](https://img.shields.io/github/stars/cloudposse/sudosh.svg)](https://github.com/cloudposse/sudosh/stargazers) 
[![GitHub Issues](https://img.shields.io/github/issues/cloudposse/sudosh.svg)](https://github.com/cloudposse/sudosh/issues)
[![Average time to resolve an issue](http://isitmaintained.com/badge/resolution/cloudposse/sudosh.svg)](http://isitmaintained.com/project/cloudposse/sudosh "Average time to resolve an issue")
[![Percentage of issues still open](http://isitmaintained.com/badge/open/cloudposse/sudosh.svg)](http://isitmaintained.com/project/cloudposse/sudosh "Percentage of issues still open")
[![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg)](https://github.com/cloudposse/sudosh/pulls)
[![License](https://img.shields.io/badge/license-APACHE%202.0%20-brightgreen.svg)](https://github.com/cloudposse/sudosh/blob/master/LICENSE)


## Purpose

The `sudo` command provides built-in session logging. Combined with [`sudoreplay`](https://www.sudo.ws/man/1.8.13/sudoreplay.man.html) it provides an easy way to review session logs on a [bastion](https://github.com/cloudposse/bastion/) host. When used as a system login shell, it will force session logging.

[Another common pattern](https://aws.amazon.com/blogs/security/how-to-record-ssh-sessions-established-through-a-bastion-host/) is to use the OpenSSH `ForceCommand` directive in `sshd_config` combined with the `script` command to log sessions. This is ineffective because the [user can easily bypass](http://serverfault.com/a/639814) it.  Using `sudosh` provides a more secure alternative that cannot be bypassed since it does not depend on `ForceCommand`.

## Usage

Here's how to use it in 3 easy steps. Checkout the [precompiled releases](https://github.com/cloudposse/sudosh/releases) if you don't want to build it yourself..

1. Enable `sudo` logging. Edit `/etc/sudoers.d/sudosh`:

    ```
    Defaults log_output
    Defaults!/usr/bin/sudoreplay !log_output
    Defaults!/sbin/reboot !log_output
    ```

2. Add this command to `/etc/shells`:

    ```
    /usr/bin/sudosh
    ```

    **Tip**: to prevent users from using other shells to login, remove those shells from `/etc/shells`.


3. Update the user `foobar` to use the `sudosh` shell.

    ```
    chsh -s /usr/bin/sudosh foobar
    echo 'foobar ALL=(foobar) ALL' > /etc/sudoers.d/sudosh.foobar
    ```


## Other Tricks

If you want to change the default shell from `bash` to something else (e.g. `zsh`), you can symlink `sudosh` to a different name. 

To change the default shell to `zsh`, you could do:

```
ln -s /usr/bin/sudosh /usr/bin/sudosh.zsh
```

Then set the user's shell to `/usr/bin/sudosh.zsh` and add the shell to `/etc/shells`.

## About


The `sudosh` utility is maintained and funded by [Cloud Posse, LLC][website]. Like it? Please let us know at <hello@cloudposse.com>

We love Open Source Software! 

See [our other projects][community]
or [hire us][hire] to help build your next cloud-platform.

  [website]: https://cloudposse.com/
  [community]: https://github.com/cloudposse/
  [hire]: https://cloudposse.com/contact/
  
### Contributors

[![Erik Osterman](http://s.gravatar.com/avatar/88c480d4f73b813904e00a5695a454cb?s=144)](https://osterman.com) 

[Erik Osterman](https://github.com/osterman) 

 
