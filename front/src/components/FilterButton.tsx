export default function FilterButton(props: {
  isPressed: boolean | undefined;
  setFilter: (arg0: any) => void;
  name: {} | null | undefined;
}): JSX.Element {
  return (
    <button
      type="button"
      className="btn toggle-btn"
      aria-pressed={props.isPressed}
      onClick={() => props.setFilter(props.name)}
    >
      <span className="visually-hidden">Show </span>
      <span>{props.name}</span>
      <span className="visually-hidden"> tasks</span>
    </button>
  );
}
