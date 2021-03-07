import React from 'react';
import './Recipe.css';

class Recipe extends React.Component {
    render() {
        return (
            <div class="recipe">
                <h4>{this.props.recipe.name}</h4>
                <ul>
                    {this.props.recipe.ingredients.map((ingredient, index) => {
                        return <li>{ingredient}</li>
                    })}
                </ul>
            </div>
        )
    }
}

export default Recipe;