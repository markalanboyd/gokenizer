-- Define a class with a table
local MyClass = {}

-- Function definition with parameters
function MyClass:new(value)
    local instance = {
        property = value,
        another_property = 10
    }
    setmetatable(instance, self)
    self.__index = self
    return instance
end

-- Method definition
function MyClass:my_function(argument)
    -- Local variable assignment
    local local_var = argument * self.another_property
    
    -- Table creation and indexing
    local my_table = { key = "value", another_key = 123 }
    local table_value = my_table.key
    
    -- Use of standard library functions
    local length = #my_table
    local concatenated = "prefix_" .. table_value
    
    -- Numeric literals
    local hex_num = 0xABC
    local bin_num = 0b1011
    local float_num = 123.456
    local sci_num = 5.67e-3
    
    -- Boolean
    local is_true = true
    local is_false = false
    
    -- Conditional statement
    if local_var > 100 then
        print("Greater than 100")
    elseif local_var < 50 then
        print("Less than 50")
    else
        print("Between 50 and 100")
    end
    
    -- Loops
    for i = 1, 10 do
        print("Loop iteration:", i)
    end
    
    -- While loop
    local counter = 0
    while counter < 5 do
        counter = counter + 1
    end
    
    -- Repeat until loop
    repeat
        counter = counter - 1
    until counter == 0
    
    -- Table iteration
    for key, value in pairs(my_table) do
        print(key, value)
    end
    
    -- Multiline string
    local multi_line = [[
    Line 1
    Line 2
    Line 3
    ]]
    
    -- Function call with an anonymous function and closures
    local function example_closure()
        local closure_var = 100
        return function()
            closure_var = closure_var + 1
            return closure_var
        end
    end
    
    local closure = example_closure()
    print(closure()) -- 101
    print(closure()) -- 102
    
    -- Return statement
    return concatenated
end

-- Instantiating a class and calling a method
local my_object = MyClass:new(5)
print(my_object:my_function(20))


-- Note: b=2 doesn't parse correctly