package com.mychess.app.engine.pieces;

import java.util.ArrayList;
import java.util.Collection;
import java.util.List;

import com.mychess.app.engine.Alliance;
import com.mychess.app.engine.board.Board;
import com.mychess.app.engine.board.BoardUtils;
import com.mychess.app.engine.board.Move;
import com.mychess.app.engine.board.Move.AttackMove;
import com.mychess.app.engine.board.Move.MajorMove;
import com.mychess.app.engine.board.Tile;

public class Bishop extends Piece {

  private final static int[] CANDIDATE_MOVE_VECTOR_COORDINATES = {-9, -7, 7, 9};

  public Bishop(final Alliance pieceAlliance, final int piecePosition) {
    super(piecePosition, pieceAlliance);
  }

  @Override
  public Collection<Move> calculateLegalMoves(final Board board) {
    final List<Move> legalMoves = new ArrayList<>();

    for (final int candidateCoordinateOffset: CANDIDATE_MOVE_VECTOR_COORDINATES) {
      int candidateDestinationCoordinate = this.piecePosition;

      while(BoardUtils.isValidTileCoordinate(candidateDestinationCoordinate)) {

        // 1列目、8列目にいるときは特定の計算がエラーになるのではじく
        if (isFirstColumnExclusion(candidateDestinationCoordinate, candidateCoordinateOffset) ||
          isEighthColumnExclusiton(candidateDestinationCoordinate, candidateCoordinateOffset)) {
            break;
          }
        candidateDestinationCoordinate += candidateCoordinateOffset;

        if (BoardUtils.isValidTileCoordinate(candidateDestinationCoordinate)) {
          final Tile candidateDestinationTile = board.getTile(candidateDestinationCoordinate);
          if(!candidateDestinationTile.isTileOccupied()) {
            legalMoves.add(new MajorMove(board, this, candidateDestinationCoordinate));
          } else {
            final Piece pieceAtDestination = candidateDestinationTile.getPiece();
            final Alliance pieceAlliance = pieceAtDestination.getPieceAlliance();
  
            if(this.pieceAlliance != pieceAlliance) {
              legalMoves.add(new AttackMove(board, this, candidateDestinationCoordinate, pieceAtDestination));
            }
            break; // 敵ピースがあるのでそれ以上先に進めない
          }
        }
      }
    }
    
    return legalMoves;
  }

  @Override
  public String toString() {
    return PieceType.BISHOP.toString();
  }
  
  private static boolean isFirstColumnExclusion(final int currentPosition, final int candidateOffset) {
    return BoardUtils.FIRST_COLUMN[currentPosition] && (candidateOffset == -9 || candidateOffset == 7);
  }

  private static boolean isEighthColumnExclusiton(final int currentPosition, final int candidateOffset) {
    return BoardUtils.EIGHTH_COLUMN[currentPosition] && (candidateOffset == -7 || candidateOffset == 9);
  }

}
