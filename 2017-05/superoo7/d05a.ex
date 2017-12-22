defmodule D05a do
  @moduledoc """
  Documentation for D05a.
  """

  def index do
    {:ok, input} = File.read("input.txt")
    input
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(&String.to_integer/1)
    |> solve
  end

  @doc """


  ## Examples

      iex> D05a.hello
      :world

  """
  def solve(input) do
    input
    |> move(&add_one/1)
  end

  #  move(input, count, indexNo, fn_one)

  # init
  def move(input, fn_one), do: move(input, 0, 0, fn_one)

  # reset index
  def move(_input, count, indexNo, _fn_one) when indexNo<0, do: count

  def move(input, count, indexNo, fn_one) do
    if indexNo > length(input) - 1 do
      count
    else
      current_offset = Enum.at(input, indexNo)
      updated_offsets = List.update_at(input, indexNo, fn_one)
      move(updated_offsets, count + 1, indexNo + current_offset, fn_one)
    end
  end


  def add_one(value), do: value + 1
end

D05a.index
|> IO.inspect

