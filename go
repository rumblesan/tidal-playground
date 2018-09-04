#!/bin/bash

if [ -n "$TMUX" ]; then
  echo "Already in an existing tmux session"
  exit 1
fi

VIM=${EDITOR}

FILE=${FILE:-"./sets/$(date +%F).tidal"}
SESSION=${SESSION:-"tidal"}

TIDAL_TEMPO_IP=${TIDAL_TEMPO_IP:-""}

args=${@:-$FILE}

# Check if tmux session "tidal" is running, attach only
# else, create new session, split windows and run processes
tmux -2 attach-session -t $SESSION || tmux -2 \
  new-session -s $SESSION \; \
  rename-window -t 1 "techno!" \; \
  split-window -v -p 10 -t $SESSION \; \
  send-keys -t 1 "$VIM $args" C-m \; \
  send-keys -t 2 "TIDAL_TEMPO_IP=$TIDAL_TEMPO_IP ./tidal" C-m \; \
  select-pane -t 1 \; \
  new-window -n "sc" \; \
  send-keys "./supercollider/run" C-m \; \
  previous-window
