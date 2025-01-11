import React from "react"
import Knight from "./Knight"
import Square from "./Square"

interface BoardProps {
    knightPosition: [number, number];
}

function renderSquare(i:number, [knightX, knightY]: [number, number]) {
    const x = i % 8;
    const y = Math.floor(i / 8);
    const isKnightHere = x === knightX && y === knightY;
    const black = (x + y) % 2 === 1;
    const piece = isKnightHere ? <Knight /> : null;

    return (
        <div key={i} style={{ width: '12.5%', height: '12.5%' }}>
            <Square black={black}>{piece}</Square>
        </div>
    );
}

const Board:React.FC<BoardProps> = ({ knightPosition = [0,0] }) => {
    const square = []
    for (let i = 0; i < 64; i++) {
        square.push(renderSquare(i, knightPosition))
    }
    
    return (
        <div className="h-screen w-screen flex flex-wrap">
            {square}
        </div>
    )
}

export default Board