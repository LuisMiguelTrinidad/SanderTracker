import { FC } from 'react';


const Header: FC = () => {
    return (
        <header className="flex justify-center py-4 space-x-8 bg-slate-900">
            <div className="flex items-center">
                <h1 className="font-Anta text-7xl">SanderTracker</h1>
            </div>
            <div>
                <img src="src/Resources/Images/Logo.png" className="h-24"/>
            </div>
        </header>
    )
}

export default Header;