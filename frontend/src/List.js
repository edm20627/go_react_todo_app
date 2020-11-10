import React from 'react';

class List extends React.Component {
  render() {
    const todo = this.props.lists.map((list) => {
      return (
        <li key={list.id}>
          {new Date(list.createdAt).getMonth() + "/" + new Date(list.createdAt).getMonth() + " " + list.text}
          <button onClick={() => this.props.handleDelete(list)}>削除</button>
        </li>
      )
    })
    return (
      <ol>
        {todo}
      </ol>
    )
  }
}

export default List;