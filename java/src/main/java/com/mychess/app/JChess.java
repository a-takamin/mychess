package com.mychess.app;

import com.mychess.app.engine.board.Board;
import com.mychess.app.gui.Table;

/**
 * Hello world!
 *
 */
public class JChess {
  public static void main(String[] args) {
    Board board = Board.createStandardBoard();
    System.out.println(board);

    Table table = new Table();
  }
}
