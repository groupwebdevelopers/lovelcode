import React from 'react'
import Header from '../../components/Header/Header'
import DesignPlans from '../../components/DesignPlans/DesignPlans'
import Portfolio from '../../components/Portfolio/Portfolio'
import Features from '../../components/Features/Features'
import SocialNetworks from '../../components/SocialNetworks/SocialNetworks'
import NewsAndArticles from '../../components/NewsAndArticles/NewsAndArticles'
import Introduction from '../../components/Introduction/Introduction'
import Customers from '../../components/Customers/Customers'
import Footer from '../../components/Footer/Footer'

export default function Home() {
    return (
        <div>
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
