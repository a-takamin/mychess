package com.mychess.app.engine;

import com.mychess.app.engine.board.BoardUtils;
import com.mychess.app.engine.player.BlackPlayer;
import com.mychess.app.engine.player.Player;
import com.mychess.app.engine.player.WhitePlayer;

public enum Alliance {
  WHITE{ 
    @Override
    public int getDirection() {
      return -1;
    }
    @Override
    public int getOppositeDirection() {
      return 1;
    }
    @Override
    public boolean isWhite() {
      return true;
    }
    @Override
    public boolean isBlack() {
      return false;
    }
    @Override
    public boolean isPawnPromotionSquare(int position) {
      return BoardUtils.FIRST_RANK[position];
    }
    @Override
    public Player choosePlayer(final WhitePlayer whitePlayer, final BlackPlayer blackPlayer) {
      return whitePlayer;
    }
  },
  BLACK {
    @Override
    public int getDirection() {
      return 1;
    }
    @Override
    public int getOppositeDirection() {
      return -1;
    }
    @Override
    public boolean isWhite() {
      return false;
    }
    @Override
    public boolean isBlack() {
      return true;
    }
    @Override
    public boolean isPawnPromotionSquare(int position) {
      return BoardUtils.EIGHTH_RANK[position];
    }
    @Override
    public Player choosePlayer(final WhitePlayer whitePlayer, final BlackPlayer blackPlayer) {
      return blackPlayer;
    }
  },;

  public abstract int getDirection();
  public abstract int getOppositeDirection();
  public abstract boolean isWhite();
  public abstract boolean isBlack();
  public abstract boolean isPawnPromotionSquare(int position);
  public abstract Player choosePlayer(WhitePlayer whitePlayer, BlackPlayer blackPlayer);
}
