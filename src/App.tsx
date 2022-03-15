import Todo from './components/Todo';
import Form from './components/Form';
import FilterButton from './components/FilterButton';
import { useEffect, useRef, useState } from 'react';
import { nanoid } from 'nanoid';

interface Task {
  id: string | undefined;
  completed: boolean | undefined;
  name: {} | null | undefined;
}
interface FilterMap {
  All: () => boolean;
  Active: (task: any) => boolean;
  Completed: (task: any) => boolean;
  [key: string]: (task: any) => boolean;
}

function usePrevious(value: number) {
  const ref = useRef<number>();
  useEffect(() => {
    ref.current = value;
  });
  return ref.current;
}
export default function App(props: { tasks: Task[] }): JSX.Element {
  const FILTER_MAP: FilterMap = {
    All: () => true,
    Active: (task) => !task.completed,
    Completed: (task) => task.completed,
  };
  const FILTER_NAMES = Object.keys(FILTER_MAP);
  const [filter, setFilter] = useState('All');
  const filterList = FILTER_NAMES.map((filterName) => (
    <FilterButton
      key={filterName}
      name={filterName}
      isPressed={filterName === filter}
      setFilter={setFilter}
    />
  ));
  const [tasks, setTasks] = useState(props.tasks);
  const taskList = tasks
    .filter(FILTER_MAP[filter])
    .map((task) => (
      <Todo
        id={task.id}
        name={task.name}
        completed={task.completed}
        key={task.id}
        toggleTaskCompleted={toggleTaskCompleted}
        deleteTask={deleteTask}
        editTask={editTask}
      />
    ));

  function addTask(name: {} | null | undefined): void {
    const newTask = { id: 'todo-' + nanoid(), name: name, completed: false };
    setTasks([...tasks, newTask]);
  }
  function toggleTaskCompleted(id: string | undefined): void {
    const updatedTasks = tasks.map((task) => {
      // if this task has the same ID as the edited task
      if (id === task.id) {
        // use object spread to make a new object
        // whose `completed` prop has been inverted
        return { ...task, completed: !task.completed };
      }
      return task;
    });
    setTasks(updatedTasks);
  }
  function deleteTask(id: string | undefined): void {
    const remainingTasks = tasks.filter((task) => id !== task.id);
    setTasks(remainingTasks);
  }
  function editTask(id: string | undefined, newName: string | undefined): void {
    const editedTaskList = tasks.map((task) => {
      // if this task has the same ID as the edited task
      if (id === task.id) {
        return { ...task, name: newName };
      }
      return task;
    });
    setTasks(editedTaskList);
  }
  const tasksNoun = taskList.length === 1 ? 'task' : 'tasks';
  const headingText = `${taskList.length} ${tasksNoun} remaining`;
  const listHeadingRef = useRef<HTMLHeadingElement>(null);
  const prevTaskLength = usePrevious(tasks.length);
  useEffect(() => {
    if (prevTaskLength) {
      if (tasks.length - prevTaskLength === -1) {
        if (!listHeadingRef.current) throw Error('null');
        listHeadingRef.current.focus();
      }
    }
  }, [tasks.length, prevTaskLength]);

  return (
    <div className="todoapp stack-large">
      <h1>TodoMatic</h1>
      <Form addTask={addTask} />
      <div className="filters btn-group stack-exception">{filterList}</div>
      <h2 id="list-heading" tabIndex={-1} ref={listHeadingRef}>
        {headingText}
      </h2>
      <ul
        role="list"
        className="todo-list stack-large stack-exception"
        aria-labelledby="list-heading"
      >
        {taskList}
      </ul>
    </div>
  );
}
