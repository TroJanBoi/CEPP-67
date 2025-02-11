"use client";

import Board from "./_components/Board";
const DndPage = () => {

    return (
        <div className="h-screen w-screen flex justify-center items-center">
            <Board knightPosition={[1, 4]}/>
        </div>
    )
}

export default DndPage