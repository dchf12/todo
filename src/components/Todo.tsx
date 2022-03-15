export default function Todo(props: {
  id: string | undefined;
  completed: boolean | undefined;
  toggleTaskCompleted: (arg0: string | undefined) => void;
  deleteTask: (arg0: string | undefined) => void;
  name: {} | null | undefined;
}): JSX.Element {
  return (
    <li className="todo stack-small">
      <div className="c-cb">
        <input
          id={props.id}
          type="checkbox"
          defaultChecked={props.completed}
          onChange={() => props.toggleTaskCompleted(props.id)}
        />

        <label className="todo-label" htmlFor={props.id}>
          {props.name}
        </label>
      </div>
      <div className="btn-group">
        <button type="button" className="btn">
          Edit <span className="visually-hidden">{props.name}</span>
        </button>
        <button
          type="button"
          className="btn btn__danger"
          onClick={() => {
            props.deleteTask(props.id);
          }}
        >
          Delete <span className="visually-hidden">{props.name}</span>
        </button>
      </div>
    </li>
  );
}
