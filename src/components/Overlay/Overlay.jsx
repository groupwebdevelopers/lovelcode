import React, { useContext } from 'react'
import MenuContext from '../../context/MenuContext'

export default function Overlay() {
    const menuContext = useContext(MenuContext)
    console.log(menuContext.menuDisplay);
    const closeMenu = () => {
        menuContext.menuDisplayHandler(false)
    }
    return (
        <div onClick={closeMenu} className={`fixed inset-0 bg-black/20 z-20 transition-all ${menuContext.menuDisplay ? 'visible opacity-100' : 'invisible opacity-0'}`}></div>
    )
}
