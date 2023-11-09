package com.mychess.app.engine.board;

import com.mychess.app.engine.pieces.Piece;

public abstract class Move {
  
  final Board board;
  final Piece movedPiece;
  final int destinationCoodinate;

  private Move(final Board board, final Piece piece, final int destinationCoodinate) {
    this.board = board;
    this.movedPiece = piece;
    this.destinationCoodinate = destinationCoodinate;
  }

  public static final class MajorMove extends Move {
    public MajorMove(final Board board, final Piece movedPiece, final int destinationCoodinate) {
      super(board, movedPiece, destinationCoodinate);
    }
  }

  public static final class AttackMove extends Move {
    final Piece attackedPiece;
    public AttackMove(final Board board, final Piece movedPiece, final int destinationCoodinate, final Piece attackedPiece) {
      super(board, movedPiece, destinationCoodinate);
      this.attackedPiece = attackedPiece;
    }
  }

  
}
