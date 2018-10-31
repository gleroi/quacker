defmodule Message do

    defmodule MessageQuacked do 
        @type t :: %MessageQuacked{ID: String.t, authorID: String.t, content: String.t}
        defstruct [:ID, :authorID, :content]
    end

    defp new_id() do
        UUID.uuid1()
    end

    @doc """
    quack a message from authorID
    """
    @spec quack(String.t, String.t) :: MessageQuacked.t
    def quack(authorID, content) do
        %MessageQuacked{ID: new_id(), authorID: authorID, content: content}
    end
end