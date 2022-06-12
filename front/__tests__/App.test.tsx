import * as React from 'react';
import { shallow } from 'enzyme';
import App from '../src/App';

test('Initial screen of Index page ', () => {
  interface Data {
    id: string | undefined;
    completed: boolean | undefined;
    name: {} | null | undefined;
  }
  const DATA: Data[] = [
    { id: 'todo-0', name: 'Eat', completed: true },
    { id: 'todo-1', name: 'Sleep', completed: false },
    { id: 'todo-2', name: 'Repeat', completed: false },
  ];
  const wrapper = shallow(<App tasks={DATA} />);

  expect(wrapper.find('h1').text()).toEqual('TodoMatic');
  expect(wrapper.find('ul').childAt(0).key()).toEqual('todo-0');
});
