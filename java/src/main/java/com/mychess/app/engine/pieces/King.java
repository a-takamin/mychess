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

public class King extends Piece {

  private final static int[] CANDIDATE_MOVE_COORDINATES = {-9,-8,-7,-1,1,7,8,9};

  public King(final Alliance pieceAlliance, final int piecePosition) {
    super(PieceType.KING, piecePosition, pieceAlliance, true);
  }

  public King(final Alliance pieceAlliance, final int piecePosition, final boolean isFirstMove) {
    super(PieceType.KING, piecePosition, pieceAlliance, isFirstMove);
  }

  @Override
  public Collection<Move> calculateLegalMoves(final Board board) {
    final List<Move> legalMoves = new ArrayList<>();
    
    for(final int currentCandidateOffset: CANDIDATE_MOVE_COORDINATES) {
      final int candidateDestinationCoordinate = this.piecePosition + currentCandidateOffset;
      
      if(isFirstColumnExclusion(this.piecePosition, currentCandidateOffset) ||
          isEighthColumnExclusiton(this.piecePosition, currentCandidateOffset)) {
        continue;
      }
      
      if (BoardUtils.isValidTileCoordinate(candidateDestinationCoordinate)) {
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
  public King movePiece(final Move move) {
    return new King(move.getMovedPiece().getPieceAlliance(), move.getDestinationCoordinate());
  }

  @Override
  public String toString() {
    return PieceType.KING.toString();
  }

  private static boolean isFirstColumnExclusion(final int currentPosition, final int candidateOffset) {
    return BoardUtils.FIRST_COLUMN[currentPosition] && (candidateOffset == -9 || candidateOffset == -1 || candidateOffset == 7);
  }

  private static boolean isEighthColumnExclusiton(final int currentPosition, final int candidateOffset) {
    return BoardUtils.EIGHTH_COLUMN[currentPosition] && (candidateOffset == -7 || candidateOffset == 1 || candidateOffset == 9);
  }
  
}
