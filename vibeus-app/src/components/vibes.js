import React, { Component } from 'react'
import Card from '@mui/material/Card'
import CardContent from '@mui/material/CardContent'
import Typography from '@mui/material/Typography';
//import CardMedia from '@mui/material/CardMedia'

class Vibes extends Component {
    constructor(props){
        super()
    }
    render() {
        let content = this.props.content
        return(
            <Card>
                <CardContent>
                    <Typography gutterBottom variant="h5" component="div">
                        {content.user_handle}
                    </Typography>
                    <Typography variant="body1" color="text.primary">
                        {content.body}
                    </Typography>
                </CardContent>
            </Card>
        )
    }
}

export default Vibes;