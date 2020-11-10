import React from 'react';
import InputField from './InputField';
import List from './List';
import http from './http';

class Todo extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      text: "",
      lists: []
    }
    this.getTodoList = this.getTodoList.bind(this)
    this.handleChange = this.handleChange.bind(this)
    this.handleSubmit = this.handleSubmit.bind(this)
    this.handleDelete = this.handleDelete.bind(this)
  }

  componentDidMount() {
    this.getTodoList();
  }

  getTodoList(){
    return http
      .get('/todo')
      .then((response) => {
        this.setState({ lists: response.data })
      })
      .catch(error => {
        console.log(error)
      })
  }

  handleChange(e){
    this.setState({ text: e.target.value })
  }

  handleSubmit(){
    if (this.state.text === "") {
      return window.alert("入力してください")
    }
    return http
      .post('/todo', {
        text: this.state.text
      })
      .then(() => {
        this.setState({ text: "" });
        this.getTodoList();
      })
      .catch(error => {
        console.log(error)
      })
  }

  handleDelete(list){
    return http
      .delete(`/todo/${list.id}`)
      .then(() =>
        this.getTodoList()
      )
      .catch(error => {
        console.log(error)
      })
  }

  render() {
    return (
      <div className="todo">
        <div className="todo-title">
          <h1>Todo</h1>
        </div>
        <InputField
          text={this.state.text}
          handleChange={this.handleChange}
          handleSubmit={this.handleSubmit}
        />
        <List
          lists={this.state.lists}
          handleDelete={this.handleDelete}
        />
      </div>
    )
  }
}

export default Todo;