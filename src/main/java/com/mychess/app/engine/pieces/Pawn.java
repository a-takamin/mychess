package com.mychess.app.engine.pieces;

import java.util.ArrayList;
import java.util.Collection;
import java.util.List;

import com.google.common.collect.ImmutableList;
import com.mychess.app.engine.Alliance;
import com.mychess.app.engine.board.Board;
import com.mychess.app.engine.board.BoardUtils;
import com.mychess.app.engine.board.Move;
import com.mychess.app.engine.board.Move.MajorMove;

public class Pawn extends Piece {

  // 8を足し引きすれば1マス動くことになる
  private final static int[] CANDIDATE_MOVE_COORDINATES = {8, 16, 7, 9};

  public Pawn(final Alliance pieceAlliance, final int piecePosition) {
    super(piecePosition, pieceAlliance);
  }

  @Override
  public Collection<Move> calculateLegalMoves(final Board board) {
    final List<Move> legalMoves = new ArrayList<>();
    for (final int currentCandidateOffset: CANDIDATE_MOVE_COORDINATES) {
      final int candidateDestinationCoordinate = this.piecePosition + (this.getPieceAlliance().getDirection() * currentCandidateOffset);
      if(!BoardUtils.isValidTileCoordinate(candidateDestinationCoordinate)) {
        continue;
      }

      if(currentCandidateOffset == 8 && !board.getTile(candidateDestinationCoordinate).isTileOccupied()){
        // TODO: more work to do here!!!
        legalMoves.add(new MajorMove(board, this, candidateDestinationCoordinate));

      } else if(currentCandidateOffset == 16 && this.isFirstMove() && 
        BoardUtils.SECOND_ROW[this.piecePosition] && this.getPieceAlliance().isBlack() || 
        BoardUtils.SEVENTH_ROW[this.piecePosition] && this.getPieceAlliance().isWhite()) {

        // 
        final int behindCandidateDestinationCoordinate = this.piecePosition + (this.pieceAlliance.getDirection() * 8);
        if (!board.getTile(behindCandidateDestinationCoordinate).isTileOccupied() && 
            !board.getTile(candidateDestinationCoordinate).isTileOccupied()) {
          legalMoves.add(new MajorMove(board, this, candidateDestinationCoordinate));
        }
      } else if (currentCandidateOffset == 7 && 
                !(BoardUtils.EIGHTH_COLUMN[this.piecePosition] && this.getPieceAlliance().isWhite() ||
                (BoardUtils.FIRST_COLUMN[this.piecePosition] && this.getPieceAlliance().isBlack()))) {
        if(board.getTile(candidateDestinationCoordinate).isTileOccupied()) {
          final Piece pieceOnCandidate = board.getTile(candidateDestinationCoordinate).getPiece();
          if(this.pieceAlliance != pieceOnCandidate.getPieceAlliance()) {
            // TODO: more to do here
            legalMoves.add(new MajorMove(board, this, candidateDestinationCoordinate));
          }
        }

      } else if (currentCandidateOffset == 9 && 
                !(BoardUtils.FIRST_COLUMN[this.piecePosition] && this.getPieceAlliance().isWhite() ||
                (BoardUtils.EIGHTH_COLUMN[this.piecePosition] && this.getPieceAlliance().isBlack()))) {
        if(board.getTile(candidateDestinationCoordinate).isTileOccupied()) {
          final Piece pieceOnCandidate = board.getTile(candidateDestinationCoordinate).getPiece();
          if(this.pieceAlliance != pieceOnCandidate.getPieceAlliance()) {
            // TODO: more to do here
            legalMoves.add(new MajorMove(board, this, candidateDestinationCoordinate));
          }
        }

      }
    }

    return ImmutableList.copyOf(legalMoves);
  }

  @Override
  public String toString() {
    return PieceType.PAWN.toString();
  }
  
}