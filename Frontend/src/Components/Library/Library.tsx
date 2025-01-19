import { FC } from 'react';

import SagaCard from './SagaCard.tsx';

import Saga from '../../Types/Saga.ts';

const Library: FC<{sagas: Saga[]}> = ({ sagas }) => {

    return (
        <div className='p-4 space-y-8'>
            {sagas.map((saga: Saga) => <SagaCard saga={saga} key={saga.id}/>)}
        </div>
    );
}

export default Library;