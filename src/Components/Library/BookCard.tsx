import { FC } from 'react';
import { useState } from 'react';

import Book from '../../Types/Book';

const BookCard: FC<{ book: Book }> = ({ book }) => {
    const [progress, setProgress] = useState<number>(0);
    const progressPercentage = (progress / book.pages) * 100;

    console.log(progressPercentage);
    
    return (
        <div className='flex flex-col space-y-4'>
            <div className='flex flex-col flex-1 p-4 space-y-4 rounded-md bg-slate-900'>
                {/*Title and subtitle*/}
                <div className='flex flex-col flex-1 space-y-2'>
                    <div className='flex items-center flex-1'>
                        <h1 className='inline-block text-3xl leading-none align-middle transition-all duration-1000 font-Anta group-hover:text-orange-300'>{book.title}</h1>
                    </div>
                    <div>
                        <h2 className='text-lg leading-none transition-all duration-1000 font-Anta group-hover:text-orange-300'>{book.author}</h2>
                    </div>
                </div>

                {/*Cover*/}
                <div className='space-y-4'>
                    <div>
                        <img className='rounded-md' src={book.cover} alt={book.title} />
                    </div>
                    <div className='flex justify-center space-x-2 font-Anta'>
                        <input className='flex-1 h-10 min-w-0 p-2 text-2xl leading-none text-center rounded-md bg-slate-700' type="number" max={book.pages} min={0} value={progress} onChange={k => setProgress(Number(k.target.value))}/>
                        <p className='flex-none text-4xl font-bold transition-all duration-1000 group-hover:text-orange-300'>/</p>
                        <input className='flex-1 h-10 min-w-0 p-2 text-2xl leading-none text-center rounded-md bg-slate-900' type="number" value={book.pages} readOnly />
                    </div>
                </div>
            </div>
            <div className='box-border flex h-10 border-4 rounded-xl'>
            <div className="h-full bg-orange-300 rounded-lg" style={{ width: `${progressPercentage}%` }}></div>
            </div>
        </div>
    )
}

export default BookCard;