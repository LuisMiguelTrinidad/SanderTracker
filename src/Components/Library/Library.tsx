import { FC } from 'react';


import SagaCard from './SagaCard.tsx';

import Saga from '../../Types/Saga.tsx';

const Library: FC = () => {

    const sagas: Saga[] = [
        {
            id: 1,
            title: 'Middle Earth saga',
            books: [
                {
                    id: '1',
                    title: 'The Hobbit',
                    author: 'J.R.R. Tolkien',
                    cover: 'https://images-na.ssl-images-amazon.com/images/I/91bYsX41DVL.jpg',
                    pages: 310,
                    isbn: '978-0618260300',
                    publishDate: new Date('09/21/1937'),
                },
                {
                    id: '2',
                    title: 'The Fellowship of the Ring',
                    author: 'J.R.R. Tolkien',
                    cover: 'https://images-na.ssl-images-amazon.com/images/I/91bYsX41DVL.jpg',
                    pages: 398,
                    isbn: '978-0618346257',
                    publishDate: new Date('07/29/1954'),
                },
                {
                    id: '3',
                    title: 'The Two Towers',
                    author: 'J.R.R. Tolkien',
                    cover: 'https://images-na.ssl-images-amazon.com/images/I/91bYsX41DVL.jpg',
                    pages: 327,
                    isbn: '978-0618346264',
                    publishDate: new Date('11/11/1954'),
                },
                {
                    id: '4',
                    title: 'The Return of the King',
                    author: 'J.R.R. Tolkien',
                    cover: 'https://images-na.ssl-images-amazon.com/images/I/91bYsX41DVL.jpg',
                    pages: 347,
                    isbn: '978-0618346271',
                    publishDate: new Date('10/20/1955'),
                },
                {
                    id: '5',
                    title: 'The Silmarillion',
                    author: 'J.R.R. Tolkien',
                    cover: 'https://images-na.ssl-images-amazon.com/images/I/91bYsX41DVL.jpg',
                    pages: 365,
                    isbn: '978-0618391110',
                    publishDate: new Date('09/15/1977'),
                }
            ]
        },
        {
            id: 2,
            title: 'Middle Earth saga',
            books: [
                {
                    id: '1',
                    title: 'The Hobbit',
                    author: 'J.R.R. Tolkien',
                    cover: 'https://images-na.ssl-images-amazon.com/images/I/91bYsX41DVL.jpg',
                    pages: 310,
                    isbn: '978-0618260300',
                    publishDate: new Date('09/21/1937'),
                },
                {
                    id: '2',
                    title: 'The Fellowship of the Ring',
                    author: 'J.R.R. Tolkien',
                    cover: 'https://images-na.ssl-images-amazon.com/images/I/91bYsX41DVL.jpg',
                    pages: 398,
                    isbn: '978-0618346257',
                    publishDate: new Date('07/29/1954'),
                },
                {
                    id: '3',
                    title: 'The Two Towers',
                    author: 'J.R.R. Tolkien',
                    cover: 'https://images-na.ssl-images-amazon.com/images/I/91bYsX41DVL.jpg',
                    pages: 327,
                    isbn: '978-0618346264',
                    publishDate: new Date('11/11/1954'),
                },
                {
                    id: '4',
                    title: 'The Return of the King',
                    author: 'J.R.R. Tolkien',
                    cover: 'https://images-na.ssl-images-amazon.com/images/I/91bYsX41DVL.jpg',
                    pages: 347,
                    isbn: '978-0618346271',
                    publishDate: new Date('10/20/1955'),
                },
                {
                    id: '5',
                    title: 'The Silmarillion',
                    author: 'J.R.R. Tolkien',
                    cover: 'https://images-na.ssl-images-amazon.com/images/I/91bYsX41DVL.jpg',
                    pages: 365,
                    isbn: '978-0618391110',
                    publishDate: new Date('09/15/1977'),
                }
            ]
        },
    ];

    return (
        <div className='p-4 space-y-8'>
            {sagas.map((saga: Saga) => <SagaCard saga={saga} key={saga.id}/>)}
        </div>
    );
}

export default Library;