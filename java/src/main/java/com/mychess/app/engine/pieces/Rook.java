package com.mychess.app.engine.pieces;

import java.util.ArrayList;
import java.util.Collection;
import java.util.List;

import com.mychess.app.engine.Alliance;
import com.mychess.app.engine.board.Board;
import com.mychess.app.engine.board.BoardUtils;
import com.mychess.app.engine.board.Move;
import com.mychess.app.engine.board.Move.MajorAttackMove;
import com.mychess.app.engine.board.Move.MajorMove;
import com.mychess.app.engine.board.Tile;

public class Rook extends Piece {

  private final static int[] CANDIDATE_MOVE_VECTOR_COORDINATES = {-8, -1, 1, 8};
  
  public Rook(final Alliance pieceAlliance, final int piecePosition) {
    super(PieceType.ROOK, piecePosition, pieceAlliance, true);
  }

  public Rook(final Alliance pieceAlliance, final int piecePosition, boolean isFirstMove) {
    super(PieceType.ROOK, piecePosition, pieceAlliance, isFirstMove);
  }
  
  @Override
  public Collection<Move> calculateLegalMoves(Board board) {
    final List<Move> legalMoves = new ArrayList<>();
    
    for(final int candidateCoordinateOffset: CANDIDATE_MOVE_VECTOR_COORDINATES) {
      int candidateDestinationCoordinate = this.piecePosition;
      while(BoardUtils.isValidTileCoordinate(candidateDestinationCoordinate)){
        if (isFirstColumnExclusion(candidateDestinationCoordinate, candidateCoordinateOffset) ||
          isEighthColumnExclusiton(candidateDestinationCoordinate, candidateCoordinateOffset)) {
          break;
        }
        candidateDestinationCoordinate += candidateCoordinateOffset;
        if(BoardUtils.isValidTileCoordinate(candidateDestinationCoordinate)) {
          final Tile candidateDestinationTile = board.getTile(candidateDestinationCoordinate);
          if(!candidateDestinationTile.isTileOccupied()) {
            legalMoves.add(new MajorMove(board, this, candidateDestinationCoordinate));
        } else {
          final Piece pieceAtDestination = candidateDestinationTile.getPiece();
            final Alliance pieceAlliance = pieceAtDestination.getPieceAlliance();
  
            if(this.pieceAlliance != pieceAlliance) {
              legalMoves.add(new MajorAttackMove(board, this, candidateDestinationCoordinate, pieceAtDestination));
            }
            break; // 敵ピースがあるのでそれ以上先に進めない
          }
        }
      }
    }
    return legalMoves;
  }

  @Override
  public Rook movePiece(final Move move) {
    return new Rook(move.getMovedPiece().getPieceAlliance(), move.getDestinationCoordinate());
  }

  @Override
  public String toString() {
    return PieceType.ROOK.toString();
  }

  private static boolean isFirstColumnExclusion(final int currentPosition, final int candidateOffset) {
    return BoardUtils.FIRST_COLUMN[currentPosition] && (candidateOffset == -1);
  }

  private static boolean isEighthColumnExclusiton(final int currentPosition, final int candidateOffset) {
    return BoardUtils.EIGHTH_COLUMN[currentPosition] && (candidateOffset == 1);
  }
  
}
