package com.mychess.app.engine.board;

import java.util.HashMap;
import java.util.Map;

import com.google.common.collect.ImmutableMap;
import com.mychess.app.engine.Alliance;
import com.mychess.app.engine.pieces.Piece;

public abstract class Tile {

  protected final int tileCoordinate; // 座標
  private static final Map<Integer, EmptyTile> EMPTY_TILES_CACHE = createAllPossibleEmtpyTiles();

  private static Map<Integer, EmptyTile> createAllPossibleEmtpyTiles(){  
    final Map<Integer, EmptyTile> emptyTileMap = new HashMap<>();
    for (int i= 0; i< BoardUtils.NUM_TILES; i++){
      emptyTileMap.put(i, new EmptyTile(i));
    }
    return ImmutableMap.copyOf(emptyTileMap);
  }
  
  public static Tile createTile(final int tileCoordinate, final Piece piece){
    return piece != null ? new OccupiedTile(tileCoordinate, piece) : EMPTY_TILES_CACHE.get(tileCoordinate);
  }

  private Tile(int tileCoordinate) {
    this.tileCoordinate = tileCoordinate;
  }

  public abstract boolean isTileOccupied();

  public abstract Piece getPiece();

  public static final class EmptyTile extends Tile{
    EmptyTile(final int coordinate) {
      super(coordinate);
    }

    @Override
    public String toString() {
      return "-";
    }

    @Override
    public boolean isTileOccupied(){
      return false;
    }

    @Override
    public Piece getPiece(){
      return null;
    }
  }

  public static final class OccupiedTile extends Tile {
    
    private final Piece pieceOnTile;
    
    OccupiedTile(int coordinate, Piece pieceOnTile) {
      super(coordinate);
      this.pieceOnTile = pieceOnTile;
    }

    @Override
    public String toString() {
      return getPiece().getPieceAlliance().isBlack() ? getPiece().toString().toLowerCase() : getPiece().toString();
    }

    @Override
    public boolean isTileOccupied(){
      return true;
    }

    @Override
    public Piece getPiece(){
      return this.pieceOnTile;
    }
  }
}