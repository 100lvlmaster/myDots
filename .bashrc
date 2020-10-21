#
# ~/.bashrc
#

# If not running interactively, don't do anything
[[ $- != *i* ]] && return
wal -qR
alias connect='adb connect 192.168.0.110:5555'
export PATH="$PATH:/opt/flutter/bin"
alias ls='ls --color=auto'
PS1='[\u@\h \W]\$ '
eval "$(starship init bash)"
