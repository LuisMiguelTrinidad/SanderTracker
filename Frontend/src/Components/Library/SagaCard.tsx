import { FC } from 'react';
import { useState } from 'react';

import Saga from "../../Types/Saga"
import Book from "../../Types/Book"

import BookCard from "./BookCard";

const SagaCard: FC<{ saga: Saga }> = ({ saga }) => {
    const [open, setOpen] = useState<boolean>(false);
    return (
        <div className="box-border p-4 space-y-8 transition-all duration-1000 border-4 rounded-md hover:border-orange-300 hover:shadow-equal-md hover:shadow-orange-300 group">
            <div className="flex items-center py-4 space-x-1" onClick={() => setOpen(!open)}>
                <svg className='w-8 h-8 transition-all duration-1000 fill-white group-hover:fill-orange-300' xmlns="http://www.w3.org/2000/svg" viewBox="0 -960 960 960">
                    {open ? (<path d="M480-560 280-360h400L480-560Z"/>) : (<path d="M480-360 280-560h400L480-360Z"/>)}
                </svg>
                <p className="text-4xl leading-none transition-all duration-1000 font-Anta group-hover:text-orange-300">{saga.title}</p>
            </div>
            <div className={`grid grid-cols-6 gap-6 ${open ? 'block' : 'hidden'}`}>
                {saga.books.map((book: Book) => <BookCard book={book} key={book.id}/>)}
            </div>
        </div>
    )
}

export default SagaCard;