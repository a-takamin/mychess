package com.mychess.app.engine.pieces;

import java.util.ArrayList;
import java.util.Collection;
import java.util.List;

import com.google.common.collect.ImmutableList;
import com.mychess.app.engine.Alliance;
import com.mychess.app.engine.board.Board;
import com.mychess.app.engine.board.BoardUtils;
import com.mychess.app.engine.board.Move;
import com.mychess.app.engine.board.Move.MajorAttackMove;
import com.mychess.app.engine.board.Move.MajorMove;
import com.mychess.app.engine.board.Tile;

public class Knight extends Piece {

  private final static int[] CANDIDATE_MOVE_COORDINATES = {-17, -15, -10, -6, 6, 10, 15, 17};

  public Knight(final Alliance pieceAlliance, final int piecePosition) {
    super(PieceType.KNIGHT, piecePosition, pieceAlliance, true);
  }

  public Knight(final Alliance pieceAlliance, final int piecePosition, final boolean isFirstMove) {
    super(PieceType.KNIGHT, piecePosition, pieceAlliance, isFirstMove);
  }

  @Override
  public Collection<Move> calculateLegalMoves(final Board board) {
  
    final List<Move> legalMoves = new ArrayList<>();

    for(final int currentCandidateOffset: CANDIDATE_MOVE_COORDINATES) {
      final int candidateDestinationCoordinate = this.piecePosition + currentCandidateOffset;
      if(BoardUtils.isValidTileCoordinate(candidateDestinationCoordinate)) {

        // 1,2,7,8列目にいるときは特定の計算がエラーになるのではじく
        if(isFirstColumnExclusion(this.piecePosition, currentCandidateOffset) ||
            isSecondColumnExclusiton(this.piecePosition, currentCandidateOffset) ||
            isSeventhColumnExclusion(this.piecePosition, currentCandidateOffset) ||
            isEighthColumnExclusiton(this.piecePosition, currentCandidateOffset)
        ){
          continue;
        }

        final Tile candidateDestinationTile = board.getTile(candidateDestinationCoordinate);
        if(!candidateDestinationTile.isTileOccupied()) {
          legalMoves.add(new MajorMove(board, this, candidateDestinationCoordinate));
        } else {
          final Piece pieceAtDestination = candidateDestinationTile.getPiece();
          final Alliance pieceAlliance = pieceAtDestination.getPieceAlliance();

          if(this.pieceAlliance != pieceAlliance) {
            legalMoves.add(new MajorAttackMove(board, this, candidateDestinationCoordinate, pieceAtDestination));
          }
        }

      }
    }
    return ImmutableList.copyOf(legalMoves);
  }

  @Override
  public Knight movePiece(final Move move) {
    return new Knight(move.getMovedPiece().getPieceAlliance(), move.getDestinationCoordinate());
  }

  @Override
  public String toString() {
    return PieceType.KNIGHT.toString();
  }

  private static boolean isFirstColumnExclusion(final int currentPosition, final int candidateOffset) {
    return BoardUtils.FIRST_COLUMN[currentPosition] && (candidateOffset == -17 || candidateOffset == -10 || candidateOffset == 6 || candidateOffset == 15);
  }

  private static boolean isSecondColumnExclusiton(final int currentPosition, final int candidateOffset) {
    return BoardUtils.SECOND_COLUMN[currentPosition] && (candidateOffset == -10 || candidateOffset == 6);
  }

  private static boolean isSeventhColumnExclusion(final int currentPosition, final int candidateOffset) {
    return BoardUtils.SEVENTH_COLUMN[currentPosition] && (candidateOffset == -6 || candidateOffset == 10);
  }

  private static boolean isEighthColumnExclusiton(final int currentPosition, final int candidateOffset) {
    return BoardUtils.EIGHTH_COLUMN[currentPosition] && (candidateOffset == -15 || candidateOffset == -6 || candidateOffset == 10 || candidateOffset == 17);
  }
    
}
