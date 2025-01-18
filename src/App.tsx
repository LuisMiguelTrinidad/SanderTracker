import { FC } from 'react';
import Header from './Components/Common/Header.jsx';
import Library from './Components/Library/Library.js';

const App: FC = () => {
  return (
    <div className='text-white bg-slate-800 min-h-dvh'>
      <Header/>
      <Library/>
    </div>
  );
}

export default App;