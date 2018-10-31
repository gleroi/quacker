defmodule QuackerTest do
  use ExUnit.Case
  doctest Quacker

  test "greets the world" do
    assert Quacker.hello() == :world
  end
end
