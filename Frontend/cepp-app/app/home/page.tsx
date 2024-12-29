"use client";

import { ReactFlow, Background, Controls } from '@xyflow/react';
import '@xyflow/react/dist/style.css';

const HomePage = () => {
    return (
        <div className='w-screen h-screen'>
            <ReactFlow>
                <Background />
                <Controls />
            </ReactFlow>
        </div>
    );
};

export default HomePage;
