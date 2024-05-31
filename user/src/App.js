import { useState } from 'react';

function Square({value, onSquareClick}) { //take value passed from Board
  // and the onSquareClick function, which updates the Board when a Square is clicked
  return ( // onClick: what happens if clicked, each square component handles click independently but the result is shared with others in Board
    <button className="square" onClick={onSquareClick}>
      {value}
    </button>
  );
}

// after 6 pieces, we don't add more pieces, only move
// move involves a first click and second click, which should be different from one click is handled. 
// the first click removes the piece at the first index, the second click adds the piece at the second index
// check if clickedOn object is the same as piece with the turn

export default function Board() {
  const [xIsNext, setXIsNext] = useState(true);
  const [squares, setSquares] = useState(Array(9).fill(null));
  const [firstClick, setFirstClick] = useState(null); // the firstClick can be anywhere, so use null to start with


  function handleClick(i) { // if occupied after 6 moves, we click to move it
    if (calculateWinner(squares)) {
      return;
    }
    const nextSquares = squares.slice();
    
    // count existing pieces
    let count_pieces = 0;
    for (let i=0; i<nextSquares.length; i++) {
      if (nextSquares[i] !== null){
        count_pieces+=1;
      }
    }

    if (count_pieces===6){
      if (firstClick===null && nextSquares[i]!==null) { // do nothing if nextsquare[i] is null, we can't move anything
        //firstClick has just been initialized to null and a piece exists at i, then we "select" this index when the first click is detected
        if ((xIsNext && nextSquares[i]==='X') || (!xIsNext && nextSquares[i]==='O')){ //validate we are moving the current piece
          setFirstClick(i);
        }
      } else if (firstClick!==null && nextSquares[i]===null){ //selected a piece in the firstClick and destination is empty, i is the destination index
        // also works if selected first piece but clicked on occupied squares, if so, firstClick is kept, and it's moved as soon as we click on an empty destination square
        // console.log(firstClick);
        if (checkValidMove(firstClick, i)){
          let selectedPiece = nextSquares[firstClick];
          if ((xIsNext && nextSquares[4]==='X') || (!xIsNext && nextSquares[4]==='O')){ // if center is occupied
            const potentialSquares = nextSquares.slice();
            potentialSquares[i] = selectedPiece;
            potentialSquares[firstClick] = null;
            if (calculateWinner(potentialSquares)) {
              // update Board if there's a winner
              nextSquares[i] = selectedPiece;
              nextSquares[firstClick] = null;
              setFirstClick(null);
              setSquares(nextSquares);
              setXIsNext(!xIsNext);
              return;
            } else if (firstClick!==4) { // occupy center piece but not winning
              // console.log("success");
              // console.log(firstClick);
              setFirstClick(null);
              // firstClick = 4; // select the center piece, if center is occupied, destination being the center would just keep firstClick the same but doesn't do anything
              // console.log(firstClick);
              // selectedPiece = nextSquares[firstClick];
              // nextSquares[i] = selectedPiece;
              // nextSquares[firstClick] = null;
            } else { // if trying to move the center piece
              nextSquares[i] = selectedPiece;
              nextSquares[firstClick] = null;
              setFirstClick(null);
              setSquares(nextSquares);
              setXIsNext(!xIsNext);
            }

          } else {

          nextSquares[i] = selectedPiece;
          nextSquares[firstClick] = null; //remove the piece at clicked square
          
          // console.log(firstClick)
          setFirstClick(null); //reset
          setSquares(nextSquares); // update squares
          setXIsNext(!xIsNext);
          }
        }
        // do nothing if moving to an occupied square (including moving to itself) or desination invalid (fails checkValidMove)
      } else if (firstClick!==null && nextSquares[i]!==null){ // nullify the second click if the destination of the first doesn't work
        setFirstClick(null);
      }
    } else { // when you have fewer than 6 pieces, can add up to 6 pieces (when we have 5 pieces can still run this one more time but after that will fail)
      if (nextSquares[i]===null){
        if (xIsNext) {
          nextSquares[i] = 'X';
        } else {
          nextSquares[i] = 'O';
        }
        setSquares(nextSquares); // update squares
        setXIsNext(!xIsNext);
      }
  }
    // console.log(count_pieces)
  }

  const winner = calculateWinner(squares);
  let status;
  if (winner) {
    status = 'Winner: ' + winner;
  } else {
    status = 'Next player: ' + (xIsNext ? 'X' : 'O');
  }

  return ( //pass value and function to each Square component
    <>
      <div className="status">{status}</div>
      <div className="board-row">
        <Square value={squares[0]} onSquareClick={() => handleClick(0)} />
        <Square value={squares[1]} onSquareClick={() => handleClick(1)} />
        <Square value={squares[2]} onSquareClick={() => handleClick(2)} />
      </div>
      <div className="board-row">
        <Square value={squares[3]} onSquareClick={() => handleClick(3)} />
        <Square value={squares[4]} onSquareClick={() => handleClick(4)} />
        <Square value={squares[5]} onSquareClick={() => handleClick(5)} />
      </div>
      <div className="board-row">
        <Square value={squares[6]} onSquareClick={() => handleClick(6)} />
        <Square value={squares[7]} onSquareClick={() => handleClick(7)} />
        <Square value={squares[8]} onSquareClick={() => handleClick(8)} />
      </div>
    </>
  );
}

function calculateWinner(squares) {
  const lines = [
    [0, 1, 2],
    [3, 4, 5],
    [6, 7, 8],
    [0, 3, 6],
    [1, 4, 7],
    [2, 5, 8],
    [0, 4, 8],
    [2, 4, 6],
  ];
  for (let i = 0; i < lines.length; i++) {
    const [a, b, c] = lines[i];
    if (squares[a] && squares[a] === squares[b] && squares[a] === squares[c]) {
      return squares[a];
    }
  }
  return null;
}

function checkValidMove(i,j) {
  const validIdx={ 
    0:[1,3,4], 
    1:[0,2,3,4,5], 
    2:[1,4,5],
    3:[0,1,4,6,7],
    4:[0,1,2,3,5,6,7,8],
    5:[1,2,4,7,8],
    6:[3,4,7],
    7:[3,4,5,6,8],
    8:[4,5,7]
};
  if (validIdx[i].includes(j)){
    return true;
  } else {
    return false;
  }
}
