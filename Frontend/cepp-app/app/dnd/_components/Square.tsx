import React from 'react';

interface SquareProps {
    black: boolean;
    children?: React.ReactNode;
}

const Square: React.FC<SquareProps> = ({ black, children }) => {

    const fill = black ? 'black' : 'white';
    const stoke = black ? 'white' : 'black';
    return (
        <div style={{ backgroundColor: fill, color: stoke, width: '100%', height: '100%' }}>
            {children}
        </div>
    );
}

export default Square;