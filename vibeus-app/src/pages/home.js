import React, { Component } from 'react'
import Grid from '@mui/material/Grid'

import Vibes from '../components/vibes'

class Home extends Component {
    constructor(props) {
        super(props)
        this.state={isLoading:true}
    }
    componentDidMount() {
        fetch('http://localhost:8090/vibes')
        .then(response => {
            if(response.ok) {
                return response.json()
            } else {
                throw new Error('Something went wrong ...');
            }
        })
        .then(data => {
            this.setState({
                vibesContent:data,
                isLoading:false
            })
        })
        .catch(error => {this.setState({error, isLoading:false})})
    }
    render() {
        const {isLoading, error, vibesContent} = this.state
        if(isLoading){
            // TODO: Build a loading component
            return(<p> Loading.. </p>)
        }
    
        if(error){
            // TODO: Build an error component
            return(<p> {this.state.error.message} </p>)
        }
        return (
            <Grid container spacing={8}>
                <Grid item xs={12} sm={6}>
                    {vibesContent.map(
                        (vibe,index) => <Vibes key={'vk_card_'+index} content={vibe}/>
                    )}
                </Grid>
                <Grid item xs={12} sm={6}>
                    <p>Profile</p>
                </Grid>
            </Grid>
        )
    }
}

export default Home;
