package com.mychess.app.engine.player;

public enum MoveStatus {

  ILLEGAL_MOVE {
    @Override
    boolean isDone() {
      return false;
    }
  },
  DONE{
    @Override
    boolean isDone() {
      return true;
    }

  }, 
  LEAVES_PLAYER_IN_CHECK{

    @Override
    boolean isDone() {
      return false;
    }
    
  };

  abstract boolean isDone();
}
