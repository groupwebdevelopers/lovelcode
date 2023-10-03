import React, { useState } from 'react'
import Menu from '../../components/Menu/Menu'
import Overlay from '../../components/Overlay/Overlay'
import Header from '../../components/Header/Header'
import DesignPlans from '../../components/DesignPlans/DesignPlans'
import Portfolio from '../../components/Portfolios/Portfolios'
import Features from '../../components/Features/Features'
import SocialNetworks from '../../components/SocialNetworks/SocialNetworks'
import NewsAndArticles from '../../components/NewsAndArticles/NewsAndArticles'
import Introduction from '../../components/Introduction/Introduction'
import Customers from '../../components/Customers/Customers'
import Footer from '../../components/Footer/Footer'

export default function Home() {
    return (
        <div className='relative'>
            <Menu />
            <Overlay />
            <Header />
            <DesignPlans />
            <Portfolio />
            <Features />
            <SocialNetworks />
            <NewsAndArticles />
            <Introduction />
            <Customers />
            <Footer />
        </div>
    )
}
