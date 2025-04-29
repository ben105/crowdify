import { useState } from 'react';

export default function DragAndDropColumns() {
  // Initial state for items in each column
  const [leftItems, setLeftItems] = useState([
    { id: 'left-1', content: 'Left Item 1', origin: 'left' },
    { id: 'left-2', content: 'Left Item 2', origin: 'left' },
    { id: 'left-3', content: 'Left Item 3', origin: 'left' },
    { id: 'left-4', content: 'Left Item 4', origin: 'left' },
  ]);
  
  const [middleItems, setMiddleItems] = useState([]);
  
  const [rightItems, setRightItems] = useState([
    { id: 'right-1', content: 'Right Item 1', origin: 'right' },
    { id: 'right-2', content: 'Right Item 2', origin: 'right' },
    { id: 'right-3', content: 'Right Item 3', origin: 'right' },
    { id: 'right-4', content: 'Right Item 4', origin: 'right' },
  ]);

  // Keep track of which item is being dragged
  const [draggedItem, setDraggedItem] = useState(null);
  const [sourceColumn, setSourceColumn] = useState(null);
  const [dropFeedback, setDropFeedback] = useState(null);

  // Handle the start of a drag operation
  const handleDragStart = (item, column) => {
    setDraggedItem(item);
    setSourceColumn(column);
    setDropFeedback(null);
  };

  // Check if the drop is allowed (left items can't go to right column)
  const isDropAllowed = (item, targetColumn) => {
    // Left items cannot be moved to the right column
    if (item.origin === 'left' && targetColumn === 'right') {
      return false;
    }
    return true;
  };

  // Handle when an item is dropped
  const handleDrop = (targetColumn) => {
    if (!draggedItem) return;
    
    // Check if the drop is allowed
    if (!isDropAllowed(draggedItem, targetColumn)) {
      setDropFeedback(`Cannot move items from the left column to the right column`);
      setDraggedItem(null);
      setSourceColumn(null);
      return;
    }

    // Remove the item from its source column
    let updatedSourceItems;
    if (sourceColumn === 'left') {
      updatedSourceItems = leftItems.filter(item => item.id !== draggedItem.id);
      setLeftItems(updatedSourceItems);
    } else if (sourceColumn === 'middle') {
      updatedSourceItems = middleItems.filter(item => item.id !== draggedItem.id);
      setMiddleItems(updatedSourceItems);
    } else if (sourceColumn === 'right') {
      updatedSourceItems = rightItems.filter(item => item.id !== draggedItem.id);
      setRightItems(updatedSourceItems);
    }

    // Add the item to the target column
    if (targetColumn === 'left') {
      setLeftItems([...leftItems, draggedItem]);
    } else if (targetColumn === 'middle') {
      setMiddleItems([...middleItems, draggedItem]);
    } else if (targetColumn === 'right') {
      setRightItems([...rightItems, draggedItem]);
    }

    // Reset drag state
    setDraggedItem(null);
    setSourceColumn(null);
    setDropFeedback(null);
  };

  // Prevent default behavior for dragover to allow drop
  const handleDragOver = (e, targetColumn) => {
    e.preventDefault();
    
    // Provide visual feedback if the drop would be invalid
    if (draggedItem && !isDropAllowed(draggedItem, targetColumn)) {
      e.dataTransfer.dropEffect = 'none';
    }
  };

  // Render a draggable item
  const renderItem = (item, column) => {
    // Add a visual indicator for items that originated from the left
    const originClass = item.origin === 'left' 
      ? 'border-l-4 border-blue-500' 
      : 'border-r-4 border-purple-500';
    
    return (
      <div
        key={item.id}
        className={`p-3 mb-2 bg-white border border-gray-300 rounded shadow-sm cursor-move ${originClass}`}
        draggable
        onDragStart={() => handleDragStart(item, column)}
      >
        {item.content}
        <div className="text-xs text-gray-500 mt-1">Origin: {item.origin}</div>
      </div>
    );
  };

  return (
    <div className="p-6 bg-gray-100 min-h-screen">
      <h1 className="text-2xl font-bold mb-6 text-center">Drag and Drop Interface</h1>
      <p className="mb-4 text-center text-gray-600">Drag items between columns (Left items cannot be moved to the right column)</p>
      
      {dropFeedback && (
        <div className="mb-4 p-2 bg-red-100 text-red-700 rounded text-center">
          {dropFeedback}
        </div>
      )}
      
      <div className="flex gap-4">
        {/* Left Column */}
        <div 
          className="flex-1 p-4 bg-blue-50 rounded-lg min-h-96 border-2 border-dashed border-blue-200"
          onDrop={() => handleDrop('left')}
          onDragOver={(e) => handleDragOver(e, 'left')}
        >
          <h2 className="text-lg font-semibold mb-4 text-center text-blue-800">Left Items</h2>
          <div className="space-y-2">
            {leftItems.map(item => renderItem(item, 'left'))}
          </div>
        </div>

        {/* Middle Column */}
        <div 
          className="flex-1 p-4 bg-green-50 rounded-lg min-h-96 border-2 border-dashed border-green-200"
          onDrop={() => handleDrop('middle')}
          onDragOver={(e) => handleDragOver(e, 'middle')}
        >
          <h2 className="text-lg font-semibold mb-4 text-center text-green-800">Middle Column</h2>
          <div className="space-y-2">
            {middleItems.map(item => renderItem(item, 'middle'))}
          </div>
        </div>

        {/* Right Column */}
        <div 
          className="flex-1 p-4 bg-purple-50 rounded-lg min-h-96 border-2 border-dashed border-purple-200"
          onDrop={() => handleDrop('right')}
          onDragOver={(e) => handleDragOver(e, 'right')}
        >
          <h2 className="text-lg font-semibold mb-4 text-center text-purple-800">Right Items</h2>
          <div className="space-y-2">
            {rightItems.map(item => renderItem(item, 'right'))}
          </div>
        </div>
      </div>
    </div>
  );
}