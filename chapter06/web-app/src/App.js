import React from 'react';
import './App.css';
import Recipe from './Recipe';
import Navbar from './Navbar';

class App extends React.Component {
  constructor(props) {
    super(props)

    this.state = {
      recipes: []
    }

    this.getRecipes();
  }

  getRecipes() {
    fetch('/api/recipes')
      .then(response => response.json())
      .then(data => this.setState({ recipes: data }));
  }

  render() {
    return (<div>
      <Navbar />
      {this.state.recipes.map((recipe, index) => (
        <Recipe recipe={recipe} />
      ))}
    </div>);
  }
}

export default App;
