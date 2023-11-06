package com.mychess.app.engine.pieces;

import java.util.Collection;

import com.mychess.app.engine.Alliance;
import com.mychess.app.engine.board.Board;
import com.mychess.app.engine.board.Move;

public abstract class Piece {
  protected final int piecePosition;
  protected final Alliance pieceAlliance;
  
  Piece(final int piecePosition, final Alliance pieceAlliance) {
    this.piecePosition = piecePosition;
    this.pieceAlliance = pieceAlliance;
  }

  public Alliance getPieceAlliance() {
    return this.pieceAlliance;
  }

  public abstract Collection<Move> calculateLegalMoves(final Board board);
}
