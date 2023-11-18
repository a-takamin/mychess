package com.mychess.app.engine.board;

import com.mychess.app.engine.board.Board.Builder;
import com.mychess.app.engine.pieces.Pawn;
import com.mychess.app.engine.pieces.Piece;
import com.mychess.app.engine.pieces.Rook;

public abstract class Move {

  protected final Board board;
  protected final Piece movedPiece;
  protected final int destinationCoodinate;
  protected final boolean isFirstMove;

  public static final Move NULL_MOVE = new NullMove();

  private Move(final Board board, final Piece piece, final int destinationCoodinate) {
    this.board = board;
    this.movedPiece = piece;
    this.destinationCoodinate = destinationCoodinate;
    this.isFirstMove = movedPiece.isFirstMove();
  }

  private Move(final Board board, final int destinationCoodinate) {
    this.board = board;
    this.movedPiece = null;
    this.destinationCoodinate = destinationCoodinate;
    this.isFirstMove = false;
  }

  @Override
  public int hashCode() {
    final int prime = 31;
    int result = 1;
    result = prime * result + this.destinationCoodinate;
    result = prime * result + this.movedPiece.hashCode();
    result = prime * result + this.movedPiece.getPiecePosition();
    return result;
  }

  @Override
  public boolean equals(final Object other) {
    if (this == other) {
      return true;
    }
    if (!(other instanceof Move)) {
      return false;
    }
    final Move otherMove = (Move) other;
    return getCurrentCoordinate() == otherMove.getCurrentCoordinate()
        && getDestinationCoordinate() == otherMove.getDestinationCoordinate()
        && getMovedPiece().equals(otherMove.getMovedPiece());
  }

  public Board getBoard() {
    return this.board;
  }

  public int getCurrentCoordinate() {
    return this.movedPiece.getPiecePosition();
  }

  public int getDestinationCoordinate() {
    return this.destinationCoodinate;
  }

  public Piece getMovedPiece() {
    return this.movedPiece;
  }

  public boolean isAttack() {
    return false;
  }

  public boolean isCastlingMove() {
    return false;
  }

  public Piece getAttackedPiece() {
    return null;
  }

  public Board execute() {
    final Board.Builder builder = new Board.Builder();
    // 動かそうとしているコマ以外のコマをすべて取り出す
    for (final Piece piece : this.board.currentPlayer().getActivePieces()) {
      // TODO: hashcode and equals for pieces
      if (!this.movedPiece.equals(piece)) {
        builder.setPiece(piece);
      }
    }

    // 相手側のコマも。
    for (final Piece piece : this.board.currentPlayer().getOpponent().getActivePieces()) {
      builder.setPiece(piece);
    }

    // 今回動かすコマ
    builder.setPiece(this.movedPiece.movePiece(this));
    // 次のターンにコマを動かす人（MoveMaker）
    builder.setMoveMaker(this.board.currentPlayer().getOpponent().getAlliance());
    // 動かした状態で新しいボードを生成（イミュータブル）
    return builder.build();
  }

  public static class MajorAttackMove extends AttackMove {

    public MajorAttackMove(final Board board, final Piece movedPiece, final int destinationCoordinate, final Piece pieceAttacked) {
      super(board, movedPiece, destinationCoordinate, pieceAttacked);
    }

    @Override
    public boolean equals(final Object other) {
      return this == other || other instanceof MajorAttackMove && super.equals(other);
    }

    @Override
    public String toString() {
      return movedPiece.getPieceType() + BoardUtils.getPositionAtCoordinate(this.destinationCoodinate);
    }
  }

  public static final class MajorMove extends Move {
    public MajorMove(final Board board, final Piece movedPiece, final int destinationCoodinate) {
      super(board, movedPiece, destinationCoodinate);
    }

    @Override
    public boolean equals(final Object other) {
      return this == other || other instanceof MajorMove && super.equals(other);
    }

    @Override
    public String toString() {
      return movedPiece.getPieceType().toString() + BoardUtils.getPositionAtCoordinate(this.destinationCoodinate);
    }

  }

  public static class AttackMove extends Move {
    final Piece attackedPiece;

    public AttackMove(final Board board, final Piece movedPiece, final int destinationCoodinate,
        final Piece attackedPiece) {
      super(board, movedPiece, destinationCoodinate);
      this.attackedPiece = attackedPiece;
    }

    @Override
    public int hashCode() {
      return this.attackedPiece.hashCode() + super.hashCode();
    }

    @Override
    public boolean equals(final Object other) {
      if (this == other) {
        return true;
      }
      if (!(other instanceof AttackMove)) {
        return false;
      }
      final AttackMove otherAttackMove = (AttackMove) other;
      return super.equals(otherAttackMove) && getAttackedPiece().equals(otherAttackMove.getAttackedPiece());
    }

    @Override
    public boolean isAttack() {
      return true;
    }

    @Override
    public Piece getAttackedPiece() {
      return this.attackedPiece;
    }
  }

  public static final class PawnMove extends Move {

    public PawnMove(final Board board, final Piece movedPiece, final int destinationCoordinate) {
      super(board, movedPiece, destinationCoordinate);
    }

    @Override
    public boolean equals(final Object other) {
      return this == other || other instanceof PawnMove && super.equals(other);
    }

    @Override
    public String toString() {
      return BoardUtils.getPositionAtCoordinate(this.destinationCoodinate);
    }
  }

  public static class PawnAttackMove extends AttackMove {
    public PawnAttackMove(final Board board, final Piece movedPiece, final int destinationCoodinate,
        final Piece attackedPiece) {
      super(board, movedPiece, destinationCoodinate, attackedPiece);
    }

    @Override
    public boolean equals(final Object other) {
      return this == other || other instanceof PawnAttackMove && super.equals(other);
    }

    @Override
    public String toString() {
      return BoardUtils.getPositionAtCoordinate(this.movedPiece.getPiecePosition()).substring(0, 1) + "x" +
          BoardUtils.getPositionAtCoordinate(this.destinationCoodinate);
    }
  }

  public static final class PawnEnPassantAttackMove extends PawnAttackMove {
    public PawnEnPassantAttackMove(final Board board, final Piece movedPiece, final int destinationCoodinate,
        final Piece attackedPiece) {
      super(board, movedPiece, destinationCoodinate, attackedPiece);
    }

    @Override
    public boolean equals(final Object other) {
      return this == other || other instanceof PawnEnPassantAttackMove && super.equals(other);
    }

    @Override
    public Board execute() {
      final Builder builder = new Builder();
      for(final Piece piece: this.board.currentPlayer().getActivePieces()) {
        if(!this.movedPiece.equals(piece)) {
          builder.setPiece(piece);
        }
      }
      for(final Piece piece: this.board.currentPlayer().getOpponent().getActivePieces()) {
        if(!piece.equals(this.getAttackedPiece())) {
          builder.setPiece(piece);
        }
      }
      builder.setPiece(this.movedPiece.movePiece(this));
      builder.setMoveMaker(this.board.currentPlayer().getOpponent().getAlliance());
      return builder.build();
    }
  }

  public static class PawnPromotion extends Move {

    final Move decoratedMove;
    final Pawn promotedPawn;

    public PawnPromotion(final Move decoratedMove) {
      super(decoratedMove.getBoard(), decoratedMove.getMovedPiece(), decoratedMove.getDestinationCoordinate());
      this.decoratedMove = decoratedMove;
      this.promotedPawn = (Pawn)decoratedMove.getMovedPiece();
    }

    @Override
    public Board execute() {
      final Board pawnMovedBoard = this.decoratedMove.execute();
      final Board.Builder builder = new Builder();
      for(final Piece piece : pawnMovedBoard.currentPlayer().getActivePieces()) {
        if(!this.promotedPawn.equals(piece)) {
          builder.setPiece(piece);
        }
      }
      for(final Piece piece : pawnMovedBoard.currentPlayer().getOpponent().getActivePieces()) {
          builder.setPiece(piece);
      }
      builder.setPiece(this.promotedPawn.getPromotionPiece().movePiece(this));
      builder.setMoveMaker(pawnMovedBoard.currentPlayer().getAlliance());
      return null;
    }

    @Override
    public boolean isAttack() {
      return this.decoratedMove.isAttack();
    }

    @Override
    public Piece getAttackedPiece() {
      return this.decoratedMove.getAttackedPiece();
    }

    @Override
    public String toString() {
      return "";
    }

    @Override
    public int hashCode() {
      return this.decoratedMove.hashCode() + (31* this.promotedPawn.hashCode());
    }
    @Override
    public boolean equals(final Object other){
      return this == other || other instanceof PawnPromotion && (super.equals(other));
    }
  }

  public static final class PawnJump extends Move {
    public PawnJump(final Board board, final Piece movedPiece, final int destinationCoodinate) {
      super(board, movedPiece, destinationCoodinate);
    }

    @Override
    public Board execute() {
      final Builder builder = new Builder();
      for (final Piece piece : this.board.currentPlayer().getActivePieces()) {
        if (!this.movedPiece.equals(piece)) {
          builder.setPiece(piece);
        }
      }
      for (final Piece piece : this.board.currentPlayer().getOpponent().getActivePieces()) {
        builder.setPiece(piece);
      }

      final Pawn movedPawn = (Pawn) this.movedPiece.movePiece(this);
      builder.setPiece(movedPawn);
      builder.setEnPassantPawn(movedPawn);
      builder.setMoveMaker(this.board.currentPlayer().getOpponent().getAlliance());
      return builder.build();
    }

    @Override
    public String toString() {
      return BoardUtils.getPositionAtCoordinate(this.destinationCoodinate);
    }
  }

  public static class CastleMove extends Move {

    protected final Rook castleRook;
    protected final int castleRookStart;
    protected final int castleRookDestination;

    public CastleMove(final Board board, final Piece movedPiece, final int destinationCoodinate, final Rook castleRook,
        final int castleRookStart, final int castleRookDestination) {
      super(board, movedPiece, destinationCoodinate);
      this.castleRook = castleRook;
      this.castleRookStart = castleRookStart;
      this.castleRookDestination = castleRookDestination;
    }

    public Rook getCastleRook() {
      return this.castleRook;
    }

    @Override
    public boolean isCastlingMove() {
      return true;
    }

    @Override
    public Board execute() {
      final Builder builder = new Builder();
      for (final Piece piece : this.board.currentPlayer().getActivePieces()) {
        if (!this.movedPiece.equals(piece) && !this.castleRook.equals(piece)) {
          builder.setPiece(piece);
        }
      }
      for (final Piece piece : this.board.currentPlayer().getOpponent().getActivePieces()) {
        builder.setPiece(piece);
      }

      builder.setPiece(this.movedPiece.movePiece(this));
      // TODO: isFirstMove should be false after castling
      builder.setPiece(new Rook(this.castleRook.getPieceAlliance(), this.castleRookDestination));
      builder.setMoveMaker(this.board.currentPlayer().getOpponent().getAlliance());
      return builder.build();
    }

    @Override
    public int hashCode() {
      final int prime = 31;
      int result = super.hashCode();
      result = prime * result + this.castleRook.hashCode();
      result = prime * result + this.castleRookDestination;
      return result;
    }

    @Override
    public boolean equals(final Object other) {
      if(this == other) {
        return true;
      }
      if(!(other instanceof CastleMove)) {
        return false;
      }
      final CastleMove otherCastleMove = (CastleMove)other;
      return super.equals(otherCastleMove) && this.castleRook.equals(otherCastleMove.getCastleRook());
    }
  }

  public static final class KingSideCastleMove extends CastleMove {
    public KingSideCastleMove(final Board board, final Piece movedPiece, final int destinationCoodinate,
        final Rook castleRook, final int castleRookStart, final int castleRookDestination) {
      super(board, movedPiece, destinationCoodinate, castleRook, castleRookStart, castleRookDestination);
    }

    @Override
    public boolean equals(final Object other) {
      return this == other && other instanceof KingSideCastleMove && super.equals(other);
    }

    @Override
    public String toString() {
      return "0-0";
    }
  }

  public static final class QueenSideCastleMove extends CastleMove {
    public QueenSideCastleMove(final Board board, final Piece movedPiece, final int destinationCoodinate,
        final Rook castleRook, final int castleRookStart, final int castleRookDestination) {
      super(board, movedPiece, destinationCoodinate, castleRook, castleRookStart, castleRookDestination);
    }

    @Override
    public boolean equals(final Object other) {
      return this == other && other instanceof QueenSideCastleMove && super.equals(other);
    }

    @Override
    public String toString() {
      return "0-0-0";
    }
  }

  public static final class NullMove extends Move {
    public NullMove() {
      super(null, 65);
    }

    @Override
    public Board execute() {
      throw new RuntimeException("cannot execute the null move!");
    }

    @Override
    public int getCurrentCoordinate() {
      return -1;
    }
  }

  public static class MoveFactory {
    private MoveFactory() {
      throw new RuntimeException("Not instantable");
    }

    public static Move createMove(final Board board, final int currentCoordinate, final int destinationCoodinate) {
      for (final Move move : board.getAllLegalMoves()) {
        if (move.getCurrentCoordinate() == currentCoordinate
            && move.getDestinationCoordinate() == destinationCoodinate) {
          return move;
        }
      }
      return NULL_MOVE;
    }
  }

}
