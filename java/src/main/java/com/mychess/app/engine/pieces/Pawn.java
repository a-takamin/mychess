package com.mychess.app.engine.pieces;

import java.util.ArrayList;
import java.util.Collection;
import java.util.List;

import com.google.common.collect.ImmutableList;
import com.mychess.app.engine.Alliance;
import com.mychess.app.engine.board.Board;
import com.mychess.app.engine.board.BoardUtils;
import com.mychess.app.engine.board.Move;
import com.mychess.app.engine.board.Move.PawnAttackMove;
import com.mychess.app.engine.board.Move.PawnEnPassantAttackMove;
import com.mychess.app.engine.board.Move.PawnJump;
import com.mychess.app.engine.board.Move.PawnMove;
import com.mychess.app.engine.board.Move.PawnPromotion;

public class Pawn extends Piece {

  // 8を足し引きすれば1マス動くことになる
  private final static int[] CANDIDATE_MOVE_COORDINATES = { 8, 16, 7, 9 };

  public Pawn(final Alliance pieceAlliance, final int piecePosition) {
    super(PieceType.PAWN, piecePosition, pieceAlliance, true);
  }

  public Pawn(final Alliance pieceAlliance, final int piecePosition, final boolean isFirstMove) {
    super(PieceType.PAWN, piecePosition, pieceAlliance, isFirstMove);
  }

  @Override
  public Collection<Move> calculateLegalMoves(final Board board) {
    final List<Move> legalMoves = new ArrayList<>();
    for (final int currentCandidateOffset : CANDIDATE_MOVE_COORDINATES) {
      final int candidateDestinationCoordinate = this.piecePosition
          + (this.getPieceAlliance().getDirection() * currentCandidateOffset);
      if (!BoardUtils.isValidTileCoordinate(candidateDestinationCoordinate)) {
        continue;
      }

      if (currentCandidateOffset == 8 && !board.getTile(candidateDestinationCoordinate).isTileOccupied()) {
        if (this.pieceAlliance.isPawnPromotionSquare(candidateDestinationCoordinate)) {
          legalMoves.add(new PawnPromotion(new PawnMove(board, this, candidateDestinationCoordinate)));
        } else {
          legalMoves.add(new PawnMove(board, this, candidateDestinationCoordinate));
        }

      } else if (currentCandidateOffset == 16 && this.isFirstMove() &&
          ((BoardUtils.SEVENTH_RANK[this.piecePosition] && this.getPieceAlliance().isBlack()) ||
              (BoardUtils.SECOND_RANK[this.piecePosition] && this.getPieceAlliance().isWhite()))) {

        //
        final int behindCandidateDestinationCoordinate = this.piecePosition + (this.pieceAlliance.getDirection() * 8);
        if (!board.getTile(behindCandidateDestinationCoordinate).isTileOccupied() &&
            !board.getTile(candidateDestinationCoordinate).isTileOccupied()) {
          legalMoves.add(new PawnJump(board, this, candidateDestinationCoordinate));
        }
      } else if (currentCandidateOffset == 7 &&
          !(BoardUtils.EIGHTH_COLUMN[this.piecePosition] && this.getPieceAlliance().isWhite() ||
              (BoardUtils.FIRST_COLUMN[this.piecePosition] && this.getPieceAlliance().isBlack()))) {
        if (board.getTile(candidateDestinationCoordinate).isTileOccupied()) {
          final Piece pieceOnCandidate = board.getTile(candidateDestinationCoordinate).getPiece();
          if (this.pieceAlliance != pieceOnCandidate.getPieceAlliance()) {
            if (this.pieceAlliance.isPawnPromotionSquare(candidateDestinationCoordinate)) {
              legalMoves.add(new PawnPromotion(new PawnAttackMove(board, this, candidateDestinationCoordinate, pieceOnCandidate)));
            } else {
              legalMoves.add(new PawnAttackMove(board, this, candidateDestinationCoordinate, pieceOnCandidate));
            }
          }
        } else if (board.getEnPassantPawn() != null) {
          if (board.getEnPassantPawn()
              .getPiecePosition() == (this.getPiecePosition() + (this.pieceAlliance.getOppositeDirection()))) {
            final Piece pieceOnCandidate = board.getEnPassantPawn();
            if (this.pieceAlliance != pieceOnCandidate.getPieceAlliance()) {
              legalMoves
                  .add(new PawnEnPassantAttackMove(board, this, candidateDestinationCoordinate, pieceOnCandidate));
            }
          }
        }

      } else if (currentCandidateOffset == 9 &&
          !(BoardUtils.FIRST_COLUMN[this.piecePosition] && this.getPieceAlliance().isWhite() ||
              (BoardUtils.EIGHTH_COLUMN[this.piecePosition] && this.getPieceAlliance().isBlack()))) {
        if (board.getTile(candidateDestinationCoordinate).isTileOccupied()) {
          final Piece pieceOnCandidate = board.getTile(candidateDestinationCoordinate).getPiece();
          if (this.pieceAlliance != pieceOnCandidate.getPieceAlliance()) {
            if (this.pieceAlliance.isPawnPromotionSquare(candidateDestinationCoordinate)) {
              legalMoves.add(new PawnPromotion(new PawnAttackMove(board, this, candidateDestinationCoordinate, pieceOnCandidate)));
            } else {
              legalMoves.add(new PawnAttackMove(board, this, candidateDestinationCoordinate, pieceOnCandidate));
            }
          }
        } else if (board.getEnPassantPawn() != null) {
          if (board.getEnPassantPawn()
              .getPiecePosition() == (this.getPiecePosition() + (this.pieceAlliance.getOppositeDirection()))) {
            final Piece pieceOnCandidate = board.getEnPassantPawn();
            if (this.pieceAlliance != pieceOnCandidate.getPieceAlliance()) {
              legalMoves
                  .add(new PawnEnPassantAttackMove(board, this, candidateDestinationCoordinate, pieceOnCandidate));
            }
          }
        }

      }
    }

    return ImmutableList.copyOf(legalMoves);
  }

  @Override
  public Pawn movePiece(final Move move) {
    return new Pawn(move.getMovedPiece().getPieceAlliance(), move.getDestinationCoordinate());
  }

  @Override
  public String toString() {
    return PieceType.PAWN.toString();
  }

  // only Queen for simplicity
  public Piece getPromotionPiece() {
    return new Queen(this.pieceAlliance, this.piecePosition, false);
  }

}
